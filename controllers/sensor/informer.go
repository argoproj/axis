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

package sensor

import (
	sensorinformers "github.com/argoproj/argo-events/pkg/client/sensor/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/tools/cache"
)

func (controller *Controller) instanceIDReq() (*labels.Requirement, error) {
	var instanceIDReq *labels.Requirement
	var err error
	var values []string

	op := selection.DoesNotExist

	if controller.Config.InstanceID != "" {
		op = selection.Equals
		values = []string{controller.Config.InstanceID}
	}

	instanceIDReq, err = labels.NewRequirement(LabelControllerInstanceID, op, values)
	if err != nil {
		return nil, err
	}
	controller.logger.WithField("instance-id", instanceIDReq.String()).Infoln("instance id requirement")
	return instanceIDReq, nil
}

// The sensor informer adds new sensors to the controller'sensor queue based on Add, Update, and Delete event handlers for the sensor resources
func (controller *Controller) newSensorInformer() (cache.SharedIndexInformer, error) {
	labelSelector, err := controller.instanceIDReq()
	if err != nil {
		return nil, err
	}

	sensorInformerFactory := sensorinformers.NewSharedInformerFactoryWithOptions(
		controller.sensorClient,
		sensorResyncPeriod,
		sensorinformers.WithNamespace(controller.Config.Namespace),
		sensorinformers.WithTweakListOptions(func(options *metav1.ListOptions) {
			options.LabelSelector = labelSelector.String()
		}),
	)
	informer := sensorInformerFactory.Argoproj().V1alpha1().Sensors().Informer()
	informer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err == nil {
					controller.queue.Add(key)
				}
			},
			UpdateFunc: func(old, new interface{}) {
				key, err := cache.MetaNamespaceKeyFunc(new)
				if err == nil {
					controller.queue.Add(key)
				}
			},
			DeleteFunc: func(obj interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
				if err == nil {
					controller.queue.Add(key)
				}
			},
		},
	)
	return informer, nil
}
