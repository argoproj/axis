/*
Copyright 2020 BlackRock, Inc.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EventBusLister helps list EventBuses.
type EventBusLister interface {
	// List lists all EventBuses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EventBus, err error)
	// EventBuses returns an object that can list and get EventBuses.
	EventBuses(namespace string) EventBusNamespaceLister
	EventBusListerExpansion
}

// eventBusLister implements the EventBusLister interface.
type eventBusLister struct {
	indexer cache.Indexer
}

// NewEventBusLister returns a new EventBusLister.
func NewEventBusLister(indexer cache.Indexer) EventBusLister {
	return &eventBusLister{indexer: indexer}
}

// List lists all EventBuses in the indexer.
func (s *eventBusLister) List(selector labels.Selector) (ret []*v1alpha1.EventBus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventBus))
	})
	return ret, err
}

// EventBuses returns an object that can list and get EventBuses.
func (s *eventBusLister) EventBuses(namespace string) EventBusNamespaceLister {
	return eventBusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventBusNamespaceLister helps list and get EventBuses.
type EventBusNamespaceLister interface {
	// List lists all EventBuses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EventBus, err error)
	// Get retrieves the EventBus from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EventBus, error)
	EventBusNamespaceListerExpansion
}

// eventBusNamespaceLister implements the EventBusNamespaceLister
// interface.
type eventBusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventBuses in the indexer for a given namespace.
func (s eventBusNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventBus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventBus))
	})
	return ret, err
}

// Get retrieves the EventBus from the indexer for a given namespace and name.
func (s eventBusNamespaceLister) Get(name string) (*v1alpha1.EventBus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventbus"), name)
	}
	return obj.(*v1alpha1.EventBus), nil
}
