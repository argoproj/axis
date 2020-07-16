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

package stripe

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhookendpoint"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/common/logging"
	"github.com/argoproj/argo-events/eventsources/common/webhook"
	"github.com/argoproj/argo-events/eventsources/sources"
	apicommon "github.com/argoproj/argo-events/pkg/apis/common"
)

// controller controls the webhook operations
var (
	controller = webhook.NewController()
)

// set up the activation and inactivation channels to control the state of routes.
func init() {
	go webhook.ProcessRouteStatus(controller)
}

// GetEventSourceName returns name of event source
func (el *EventListener) GetEventSourceName() string {
	return el.EventSourceName
}

// GetEventName returns name of event
func (el *EventListener) GetEventName() string {
	return el.EventName
}

// GetEventSourceType return type of event server
func (el *EventListener) GetEventSourceType() apicommon.EventSourceType {
	return apicommon.StripeEvent
}

// Implement Router
// 1. GetRoute
// 2. HandleRoute
// 3. PostActivate
// 4. PostDeactivate

// GetRoute returns the route
func (rc *Router) GetRoute() *webhook.Route {
	return rc.route
}

// HandleRoute handles incoming requests on the route
func (rc *Router) HandleRoute(writer http.ResponseWriter, request *http.Request) {
	route := rc.route

	logger := route.Logger.WithFields(
		logrus.Fields{
			logging.LabelEventSourceName: route.EventSourceName,
			logging.LabelEventName:       route.EventName,
			logging.LabelEndpoint:        route.Context.Endpoint,
			logging.LabelPort:            route.Context.Port,
			logging.LabelHTTPMethod:      route.Context.Method,
		})

	logger.Info("request a received, processing it...")

	if !route.Active {
		logger.Warn("endpoint is not active, won't process it")
		common.SendErrorResponse(writer, "endpoint is inactive")
		return
	}

	const MaxBodyBytes = int64(65536)
	request.Body = http.MaxBytesReader(writer, request.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logger.WithError(err).Errorln("error reading request body")
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var event *stripe.Event
	if err := json.Unmarshal(payload, &event); err != nil {
		logger.WithError(err).Errorln("failed to parse request body")
		common.SendErrorResponse(writer, "failed to parse the event")
		return
	}

	ok := filterEvent(event, rc.stripeEventSource.EventFilter)
	if !ok {
		logger.WithField("event-type", event.Type).Warnln("failed to pass the filters")
		common.SendSuccessResponse(writer, "invalid event")
		return
	}

	data, err := json.Marshal(event)
	if err != nil {
		logger.WithField("event-id", event.ID).Warnln("failed to marshal event into gateway response")
		common.SendSuccessResponse(writer, "invalid event")
		return
	}

	logger.Infoln("dispatching event on route's data channel...")
	route.DataCh <- data
	logger.Info("request successfully processed")
	common.SendSuccessResponse(writer, "success")
}

// PostActivate performs operations once the route is activated and ready to consume requests
func (rc *Router) PostActivate() error {
	if rc.stripeEventSource.CreateWebhook {
		route := rc.route
		stripeEventSource := rc.stripeEventSource
		logger := route.Logger.WithFields(
			logrus.Fields{
				logging.LabelEventSourceName: route.EventSourceName,
				logging.LabelEventName:       route.EventName,
				logging.LabelEndpoint:        route.Context.Endpoint,
				logging.LabelHTTPMethod:      route.Context.Method,
			})
		logger.Infoln("registering a new webhook")

		apiKey, ok := common.GetEnvFromSecret(stripeEventSource.APIKey)
		if !ok {
			return errors.New("APIKey not found in ENV")
		}

		stripe.Key = apiKey

		params := &stripe.WebhookEndpointParams{
			URL: stripe.String(common.FormattedURL(stripeEventSource.Webhook.URL, stripeEventSource.Webhook.Endpoint)),
		}
		if stripeEventSource.EventFilter != nil {
			params.EnabledEvents = stripe.StringSlice(stripeEventSource.EventFilter)
		}

		endpoint, err := webhookendpoint.New(params)
		if err != nil {
			return err
		}

		logger.WithField("endpoint-id", endpoint.ID).Infoln("new stripe webhook endpoint created")
	}
	return nil
}

// PostInactivate performs operations after the route is inactivated
func (rc *Router) PostInactivate() error {
	return nil
}

func filterEvent(event *stripe.Event, filters []string) bool {
	if filters == nil {
		return true
	}
	for _, filter := range filters {
		if event.Type == filter {
			return true
		}
	}
	return false
}

// StartListening starts an SNS event source
func (el *EventListener) StartListening(ctx context.Context, dispatch func([]byte) error) error {
	logger := logging.FromContext(ctx)
	log := logging.FromContext(ctx).WithFields(map[string]interface{}{
		logging.LabelEventSourceType: el.GetEventSourceType(),
		logging.LabelEventSourceName: el.GetEventSourceName(),
		logging.LabelEventName:       el.GetEventName(),
	})
	log.Infoln("started processing the Stripe event source...")
	defer sources.Recover(el.GetEventName())

	stripeEventSource := &el.StripeEventSource
	route := webhook.NewRoute(stripeEventSource.Webhook, logger, el.GetEventSourceName(), el.GetEventName())

	return webhook.ManageRoute(ctx, &Router{
		route:             route,
		stripeEventSource: stripeEventSource,
	}, controller, dispatch)
}
