/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/gateways"
	"github.com/argoproj/argo-events/gateways/core"
	"github.com/ghodss/yaml"
	"go.uber.org/atomic"
	"io/ioutil"
	"net/http"
	"sync"
	"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1"
)

var (
	// whether http server has started or not
	hasServerStarted atomic.Bool

	// mutex synchronizes activeRoutes
	mutex sync.Mutex
	// as http package does not provide method for unregistering routes,
	// this keeps track of configured http routes and their methods.
	// keeps endpoints as keys and corresponding http methods as a map
	activeRoutes = make(map[string]map[string]struct{})

	// gatewayConfig provides a generic configuration for a gateway
	gatewayConfig = gateways.NewGatewayConfiguration()
)

// webhook is a general purpose REST API
type webhook struct {
	// REST API endpoint
	Endpoint string

	// Method is HTTP request method that indicates the desired action to be performed for a given resource.
	// See RFC7231 Hypertext Transfer Protocol (HTTP/1.1): Semantics and Content
	Method string

	// Port on which HTTP server is listening for incoming events.
	Port string
}

// Runs a gateway configuration
func configRunner(config *gateways.ConfigContext) error {
	var err error
	var errMessage string

	// mark final gateway state
	defer gatewayConfig.GatewayCleanup(config, errMessage, err)

	gatewayConfig.Log.Info().Str("config-name", config.Data.Src).Msg("parsing configuration...")

	var h *webhook
	err = yaml.Unmarshal([]byte(config.Data.Config), &h)
	if err != nil {
		errMessage = "failed to parse configuration"
		return err
	}
	gatewayConfig.Log.Info().Interface("config", config.Data.Config).Interface("webhook", h).Msg("configuring...")

	// start a http server only if given configuration contains port information and no other
	// configuration previously started the server
	if h.Port != "" && !hasServerStarted.Load() {
		// mark http server as started
		hasServerStarted.Store(true)
		go func() {
			gatewayConfig.Log.Info().Str("http-port", h.Port).Msg("http server started listening...")
			err = http.ListenAndServe(":"+fmt.Sprintf("%s", h.Port), nil)
			if err != nil {
				errMessage = "http server stopped"
			}
			if config.Active == true {
				config.StopCh <- struct{}{}
			}
			return
		}()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// waits till disconnection from client. perform cleanup.
	go func() {
		<-config.StopCh
		config.Active = false
		gatewayConfig.Log.Info().Str("config-key", config.Data.Src).Msg("stopping the configuration...")
		// remove the endpoint and http method configuration.
		mutex.Lock()
		activeHTTPMethods := activeRoutes[h.Endpoint]
		delete(activeHTTPMethods, h.Method)
		mutex.Unlock()

		wg.Done()
	}()

	config.Active = true

	event := gatewayConfig.K8Event("configuration running", v1alpha1.NodePhaseRunning, config.Data.Src)
	err = gatewayConfig.CreateK8Event(event)
	if err != nil {
		gatewayConfig.Log.Error().Str("config-key", config.Data.Src).Err(err).Msg("failed to mark configuration as running")
		return err
	}

	// configure endpoint and http method
	if h.Endpoint != "" && h.Method != "" {
		if _, ok := activeRoutes[h.Endpoint]; !ok {
			mutex.Lock()
			activeRoutes[h.Endpoint] = make(map[string]struct{})
			// save event channel for this connection/configuration
			activeRoutes[h.Endpoint][h.Method] = struct{}{}
			mutex.Unlock()

			// add a handler for endpoint if not already added.
			http.HandleFunc(h.Endpoint, func(writer http.ResponseWriter, request *http.Request) {
				// check if http methods match and route and http method is registered.
				if _, ok := activeRoutes[h.Endpoint]; ok {
					if _, isActive := activeRoutes[h.Endpoint][request.Method]; isActive {
						gatewayConfig.Log.Info().Str("endpoint", h.Endpoint).Str("http-method", h.Method).Msg("received a request")
						body, err := ioutil.ReadAll(request.Body)
						if err != nil {
							gatewayConfig.Log.Error().Err(err).Msg("failed to parse request body")
							common.SendErrorResponse(writer)
						} else {
							gatewayConfig.Log.Info().Str("endpoint", h.Endpoint).Str("http-method", h.Method).Msg("dispatching event to gateway-processor")
							common.SendSuccessResponse(writer)
							// dispatch event to gateway transformer
							gatewayConfig.DispatchEvent(&gateways.GatewayEvent{
								Src:     config.Data.Src,
								Payload: body,
							})
						}
					} else {
						gatewayConfig.Log.Warn().Str("endpoint", h.Endpoint).Str("http-method", request.Method).Msg("endpoint and http method is not an active route")
						common.SendErrorResponse(writer)
					}
				} else {
					gatewayConfig.Log.Warn().Str("endpoint", h.Endpoint).Msg("endpoint is not active")
					common.SendErrorResponse(writer)
				}
			})
		} else {
			mutex.Lock()
			activeRoutes[h.Endpoint][h.Method] = struct{}{}
			mutex.Unlock()
		}

		gatewayConfig.Log.Info().Str("config-name", config.Data.Src).Msg("configuration is running...")
		wg.Wait()
	}
	return nil
}

func main() {
	_, err := gatewayConfig.WatchGatewayEvents(context.Background())
	if err != nil {
		gatewayConfig.Log.Panic().Err(err).Msg("failed to watch k8 events for gateway state updates")
	}
	_, err = gatewayConfig.WatchGatewayConfigMap(context.Background(), configRunner, core.ConfigDeactivator)
	if err != nil {
		gatewayConfig.Log.Panic().Err(err).Msg("failed to watch gateway configuration updates")
	}
	select {}
}
