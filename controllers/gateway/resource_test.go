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

package gateway

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/argoproj/argo-events/common"
	apicommon "github.com/argoproj/argo-events/pkg/apis/common"
	eventbusv1alpha1 "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1"
	"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1"
	eventbusclientset "github.com/argoproj/argo-events/pkg/client/eventbus/clientset/versioned"
)

var testEventBus = &eventbusv1alpha1.EventBus{
	TypeMeta: metav1.TypeMeta{
		APIVersion: eventbusv1alpha1.SchemeGroupVersion.String(),
		Kind:       "EventBus",
	},
	ObjectMeta: metav1.ObjectMeta{
		Namespace: common.DefaultControllerNamespace,
		Name:      "default",
	},
	Spec: eventbusv1alpha1.EventBusSpec{
		NATS: &eventbusv1alpha1.NATSBus{
			Native: &eventbusv1alpha1.NativeStrategy{
				Replicas: 3,
				Auth:     &eventbusv1alpha1.AuthStrategyToken,
			},
		},
	},
}

var gatewayObj = &v1alpha1.Gateway{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "fake-gateway-1",
		Namespace: common.DefaultControllerNamespace,
	},
	Spec: v1alpha1.GatewaySpec{
		EventSourceRef: &v1alpha1.EventSourceRef{
			Name: "fake-event-source",
		},
		Replica:       1,
		Type:          apicommon.WebhookEvent,
		ProcessorPort: "8080",
		Template: v1alpha1.Template{
			ServiceAccountName: "fake-sa",
			Container: &corev1.Container{
				Image: "argoproj/fake-image",
			},
		},
		Service: &v1alpha1.Service{
			Ports: []corev1.ServicePort{
				{
					Name:       "server-port",
					Port:       12000,
					TargetPort: intstr.FromInt(12000),
				},
			},
		},
	},
}

var gatewayObjNoTemplate = &v1alpha1.Gateway{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "fake-gateway-2",
		Namespace: common.DefaultControllerNamespace,
	},
	Spec: v1alpha1.GatewaySpec{
		EventSourceRef: &v1alpha1.EventSourceRef{
			Name: "fake-event-source",
		},
		Replica:       1,
		Type:          apicommon.WebhookEvent,
		ProcessorPort: "8080",
		Service: &v1alpha1.Service{
			Ports: []corev1.ServicePort{
				{
					Name:       "server-port",
					Port:       12000,
					TargetPort: intstr.FromInt(12000),
				},
			},
		},
	},
}

func TestResource_BuildServiceResource(t *testing.T) {
	gwObjs := []*v1alpha1.Gateway{gatewayObj, gatewayObjNoTemplate}
	controller := newController()
	_, err := createEventBus(controller.eventBusClient)
	assert.NoError(t, err)
	for _, gwObj := range gwObjs {
		opCtx := newGatewayContext(gwObj, controller)

		service, err := opCtx.buildServiceResource()
		assert.Nil(t, err)
		assert.NotNil(t, service)
		assert.Equal(t, service.Name, gwObj.Name+"-gateway")
		assert.Equal(t, service.Namespace, opCtx.gateway.Namespace)

		newSvc, err := controller.k8sClient.CoreV1().Services(service.Namespace).Create(service)
		assert.Nil(t, err)
		assert.NotNil(t, newSvc)
		assert.Equal(t, newSvc.Name, service.Name)
		assert.Equal(t, len(newSvc.Spec.Ports), 1)
		assert.Equal(t, newSvc.Spec.Type, corev1.ServiceTypeClusterIP)
	}
}

func TestResource_BuildDeploymentResource(t *testing.T) {
	controller := newController()
	_, err := createEventBus(controller.eventBusClient)
	assert.NoError(t, err)
	ctx := newGatewayContext(gatewayObjNoTemplate, controller)
	deployment, err := ctx.buildDeploymentResource()
	assert.Nil(t, err)
	assert.NotNil(t, deployment)

	for _, container := range deployment.Spec.Template.Spec.Containers {
		assert.NotNil(t, container.Env)
		assert.Equal(t, container.Env[0].Name, common.EnvVarNamespace)
		assert.Equal(t, container.Env[0].Value, ctx.gateway.Namespace)
		assert.Equal(t, container.Env[1].Name, common.EnvVarEventSource)
		assert.Equal(t, container.Env[1].Value, ctx.gateway.Spec.EventSourceRef.Name)
		assert.Equal(t, container.Env[2].Name, common.EnvVarResourceName)
		assert.Equal(t, container.Env[2].Value, ctx.gateway.Name)
		assert.Equal(t, container.Env[3].Name, common.EnvVarControllerInstanceID)
		assert.Equal(t, container.Env[3].Value, ctx.controller.Config.InstanceID)
		assert.Equal(t, container.Env[4].Name, common.EnvVarGatewayServerPort)
		assert.Equal(t, container.Env[4].Value, ctx.gateway.Spec.ProcessorPort)
	}

	newDeployment, err := controller.k8sClient.AppsV1().Deployments(deployment.Namespace).Create(deployment)
	assert.NoError(t, err)
	assert.NotNil(t, newDeployment)
	assert.Equal(t, newDeployment.Labels[common.LabelOwnerName], ctx.gateway.Name)
	assert.NotNil(t, newDeployment.Annotations[common.AnnotationResourceSpecHash])
}

func TestResource_CreateGatewayResourceNoTemplate(t *testing.T) {
	tests := []struct {
		name       string
		updateFunc func(ctx *gatewayContext)
		testFunc   func(controller *Controller, ctx *gatewayContext, t *testing.T)
	}{
		{
			name:       "gateway without deployment and service",
			updateFunc: func(ctx *gatewayContext) {},
			testFunc: func(controller *Controller, ctx *gatewayContext, t *testing.T) {
				deploymentMetadata := ctx.gateway.Status.Resources.Deployment
				serviceMetadata := ctx.gateway.Status.Resources.Service
				deployment, err := controller.k8sClient.AppsV1().Deployments(deploymentMetadata.Namespace).Get(deploymentMetadata.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, deployment)
				service, err := controller.k8sClient.CoreV1().Services(serviceMetadata.Namespace).Get(serviceMetadata.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, service)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controller := newController()
			createEventBus(controller.eventBusClient)
			ctx := newGatewayContext(gatewayObjNoTemplate.DeepCopy(), controller)
			test.updateFunc(ctx)
			err := ctx.createGatewayResources()
			assert.Nil(t, err)
			test.testFunc(controller, ctx, t)
		})
	}
}

func TestResource_CreateGatewayResource(t *testing.T) {
	tests := []struct {
		name       string
		updateFunc func(ctx *gatewayContext)
		testFunc   func(controller *Controller, ctx *gatewayContext, t *testing.T)
	}{
		{
			name:       "gateway with deployment and service",
			updateFunc: func(ctx *gatewayContext) {},
			testFunc: func(controller *Controller, ctx *gatewayContext, t *testing.T) {
				deploymentMetadata := ctx.gateway.Status.Resources.Deployment
				serviceMetadata := ctx.gateway.Status.Resources.Service
				deployment, err := controller.k8sClient.AppsV1().Deployments(deploymentMetadata.Namespace).Get(deploymentMetadata.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, deployment)
				service, err := controller.k8sClient.CoreV1().Services(serviceMetadata.Namespace).Get(serviceMetadata.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, service)
			},
		},
		{
			name: "gateway with zero deployment replica",
			updateFunc: func(ctx *gatewayContext) {
				ctx.gateway.Spec.Replica = 0
			},
			testFunc: func(controller *Controller, ctx *gatewayContext, t *testing.T) {
				deploymentMetadata := ctx.gateway.Status.Resources.Deployment
				deployment, err := controller.k8sClient.AppsV1().Deployments(deploymentMetadata.Namespace).Get(deploymentMetadata.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, deployment)
				assert.Equal(t, *deployment.Spec.Replicas, int32(1))
			},
		},
		{
			name: "gateway with empty service template",
			updateFunc: func(ctx *gatewayContext) {
				ctx.gateway.Spec.Service = nil
			},
			testFunc: func(controller *Controller, ctx *gatewayContext, t *testing.T) {
				deploymentMetadata := ctx.gateway.Status.Resources.Deployment
				deployment, err := controller.k8sClient.AppsV1().Deployments(deploymentMetadata.Namespace).Get(deploymentMetadata.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, deployment)
				assert.Nil(t, ctx.gateway.Status.Resources.Service)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controller := newController()
			createEventBus(controller.eventBusClient)
			ctx := newGatewayContext(gatewayObj.DeepCopy(), controller)
			test.updateFunc(ctx)
			err := ctx.createGatewayResources()
			assert.Nil(t, err)
			test.testFunc(controller, ctx, t)
		})
	}
}

func TestResource_UpdateGatewayResource(t *testing.T) {
	controller := newController()
	createEventBus(controller.eventBusClient)
	ctx := newGatewayContext(gatewayObj.DeepCopy(), controller)
	err := ctx.createGatewayResources()
	assert.Nil(t, err)

	tests := []struct {
		name       string
		updateFunc func()
		testFunc   func(t *testing.T, oldMetadata *v1alpha1.GatewayResource)
	}{
		{
			name: "update deployment resource on gateway change",
			updateFunc: func() {
				ctx.gateway.Spec.Template.ServiceAccountName = "new-sa"
			},
			testFunc: func(t *testing.T, oldMetadata *v1alpha1.GatewayResource) {
				currentMetadata := ctx.gateway.Status.Resources
				deployment, err := controller.k8sClient.AppsV1().Deployments(currentMetadata.Deployment.Namespace).Get(currentMetadata.Deployment.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, deployment)
				assert.NotEqual(t, deployment.Annotations[common.AnnotationResourceSpecHash], oldMetadata.Deployment.Annotations[common.AnnotationResourceSpecHash])
				service, err := controller.k8sClient.CoreV1().Services(currentMetadata.Service.Namespace).Get(currentMetadata.Service.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, service)
				assert.Equal(t, service.Annotations[common.AnnotationResourceSpecHash], oldMetadata.Service.Annotations[common.AnnotationResourceSpecHash])
			},
		},
		{
			name: "service deleted if gateway service spec is removed",
			updateFunc: func() {
				ctx.gateway.Spec.Service = nil
			},
			testFunc: func(t *testing.T, oldMetadata *v1alpha1.GatewayResource) {
				currentMetadata := ctx.gateway.Status.Resources
				deployment, err := controller.k8sClient.AppsV1().Deployments(currentMetadata.Deployment.Namespace).Get(currentMetadata.Deployment.Name, metav1.GetOptions{})
				assert.Nil(t, err)
				assert.NotNil(t, deployment)
				assert.Nil(t, ctx.gateway.Status.Resources.Service)
				_, err = controller.k8sClient.CoreV1().Services(oldMetadata.Service.Namespace).Get(oldMetadata.Service.Name, metav1.GetOptions{})
				assert.NotNil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			metadata := ctx.gateway.Status.Resources.DeepCopy()
			test.updateFunc()
			err := ctx.updateGatewayResources()
			assert.Nil(t, err)
			test.testFunc(t, metadata)
		})
	}
}

func createEventBus(eventBusClient eventbusclientset.Interface) (*eventbusv1alpha1.EventBus, error) {
	obj := testEventBus.DeepCopyObject()
	busCopy, ok := obj.(*eventbusv1alpha1.EventBus)
	busCopy.Status.MarkDeployed("", "")
	busCopy.Status.MarkConfigured()
	clusterID := "test"
	busCopy.Status.Config = eventbusv1alpha1.BusConfig{
		NATS: &eventbusv1alpha1.NATSConfig{
			URL:       "nats://xxxx",
			ClusterID: &clusterID,
			Auth:      &eventbusv1alpha1.AuthStrategyToken,
		},
	}
	if !ok {
		return nil, errors.New("convert error")
	}
	eb, err := eventBusClient.ArgoprojV1alpha1().EventBus(common.DefaultControllerNamespace).Create(busCopy)
	return eb, err
}
