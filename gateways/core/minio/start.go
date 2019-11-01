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

package minio

import (
	"encoding/json"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/gateways"
	apicommon "github.com/argoproj/argo-events/pkg/apis/common"
	"github.com/argoproj/argo-events/store"
	"github.com/ghodss/yaml"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

// MinioEventSourceListener implements Eventing for minio event sources
type EventListener struct {
	// logger
	logger *logrus.Logger
	// k8sClient is kubernetes client
	k8sClient kubernetes.Interface
	// namespace where gateway is deployed
	namespace string
}

// StartEventSource activates an event source and streams back events
func (listener *EventListener) StartEventSource(eventSource *gateways.EventSource, eventStream gateways.Eventing_StartEventSourceServer) error {
	listener.logger.WithField(common.LabelEventSource, eventSource.Name).Infoln("activating the event source...")

	dataCh := make(chan []byte)
	errorCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	go listener.listenEvents(eventSource, dataCh, errorCh, doneCh)
	return gateways.HandleEventsFromEventSource(eventSource.Name, eventStream, dataCh, errorCh, doneCh, listener.logger)
}

// listenEvents listens to minio bucket notifications
func (listener *EventListener) listenEvents(eventSource *gateways.EventSource, dataCh chan []byte, errorCh chan error, doneCh chan struct{}) {
	defer gateways.Recover(eventSource.Name)

	logger := listener.logger.WithField(common.LabelEventSource, eventSource.Name)

	logger.Infoln("parsing minio event source...")

	var minioEventSource *apicommon.S3Artifact
	err := yaml.Unmarshal(eventSource.Value, &minioEventSource)
	if err != nil {
		errorCh <- err
		return
	}

	logger.Info("started processing the event source...")

	logger.Info("retrieving access and secret key...")
	accessKey, err := store.GetSecrets(listener.k8sClient, listener.namespace, minioEventSource.AccessKey.Name, minioEventSource.AccessKey.Key)
	if err != nil {
		errorCh <- err
		return
	}
	secretKey, err := store.GetSecrets(listener.k8sClient, listener.namespace, minioEventSource.SecretKey.Name, minioEventSource.SecretKey.Key)
	if err != nil {
		errorCh <- err
		return
	}

	logger.Infoln("setting up a minio client...")
	minioClient, err := minio.New(minioEventSource.Endpoint, accessKey, secretKey, !minioEventSource.Insecure)
	if err != nil {
		errorCh <- err
		return
	}

	logger.Info("started listening to bucket notifications...")
	for notification := range minioClient.ListenBucketNotification(minioEventSource.Bucket.Name, minioEventSource.Filter.Prefix, minioEventSource.Filter.Suffix, minioEventSource.Events, doneCh) {
		if notification.Err != nil {
			errorCh <- notification.Err
			return
		}

		logger.Infoln("parsing notification from minio...")
		payload, err := json.Marshal(notification.Records[0])
		if err != nil {
			errorCh <- err
			return
		}

		logger.Infoln("dispatching notification on data channel...")
		dataCh <- payload
	}
}
