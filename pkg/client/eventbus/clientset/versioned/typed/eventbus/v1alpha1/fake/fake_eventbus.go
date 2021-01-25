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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventBus implements EventBusInterface
type FakeEventBus struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var eventbusResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "eventbus"}

var eventbusKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "EventBus"}

// Get takes name of the eventBus, and returns the corresponding eventBus object, and an error if there is any.
func (c *FakeEventBus) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventbusResource, c.ns, name), &v1alpha1.EventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventBus), err
}

// List takes label and field selectors, and returns the list of EventBus that match those selectors.
func (c *FakeEventBus) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventBusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventbusResource, eventbusKind, c.ns, opts), &v1alpha1.EventBusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventBusList{ListMeta: obj.(*v1alpha1.EventBusList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventBusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventBus.
func (c *FakeEventBus) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventbusResource, c.ns, opts))

}

// Create takes the representation of a eventBus and creates it.  Returns the server's representation of the eventBus, and an error, if there is any.
func (c *FakeEventBus) Create(ctx context.Context, eventBus *v1alpha1.EventBus, opts v1.CreateOptions) (result *v1alpha1.EventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventbusResource, c.ns, eventBus), &v1alpha1.EventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventBus), err
}

// Update takes the representation of a eventBus and updates it. Returns the server's representation of the eventBus, and an error, if there is any.
func (c *FakeEventBus) Update(ctx context.Context, eventBus *v1alpha1.EventBus, opts v1.UpdateOptions) (result *v1alpha1.EventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventbusResource, c.ns, eventBus), &v1alpha1.EventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventBus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventBus) UpdateStatus(ctx context.Context, eventBus *v1alpha1.EventBus, opts v1.UpdateOptions) (*v1alpha1.EventBus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventbusResource, "status", c.ns, eventBus), &v1alpha1.EventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventBus), err
}

// Delete takes name of the eventBus and deletes it. Returns an error if one occurs.
func (c *FakeEventBus) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventbusResource, c.ns, name), &v1alpha1.EventBus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventBus) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventbusResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventBusList{})
	return err
}

// Patch applies the patch and returns the patched eventBus.
func (c *FakeEventBus) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventBus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventbusResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventBus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventBus), err
}
