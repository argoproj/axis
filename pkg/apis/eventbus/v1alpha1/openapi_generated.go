// +build !ignore_autogenerated

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.BusConfig":           schema_pkg_apis_eventbus_v1alpha1_BusConfig(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBus":            schema_pkg_apis_eventbus_v1alpha1_EventBus(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBusList":        schema_pkg_apis_eventbus_v1alpha1_EventBusList(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBusSpec":        schema_pkg_apis_eventbus_v1alpha1_EventBusSpec(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBusStatus":      schema_pkg_apis_eventbus_v1alpha1_EventBusStatus(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSBus":             schema_pkg_apis_eventbus_v1alpha1_NATSBus(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSConfig":          schema_pkg_apis_eventbus_v1alpha1_NATSConfig(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NativeStrategy":      schema_pkg_apis_eventbus_v1alpha1_NativeStrategy(ref),
		"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.PersistenceStrategy": schema_pkg_apis_eventbus_v1alpha1_PersistenceStrategy(ref),
	}
}

func schema_pkg_apis_eventbus_v1alpha1_BusConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BusConfig has the finalized configuration for EventBus",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nats": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSConfig"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_EventBus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBus is the definition of a eventbus resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBusSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBusStatus"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBusSpec", "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBusStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_EventBusList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBusList is the list of eventbus resources",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "eventbus",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.EventBus", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_EventBusSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBusSpec refers to specification of eventbus resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nats": {
						SchemaProps: spec.SchemaProps{
							Description: "NATS eventbus",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSBus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSBus"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_EventBusStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBusStatus holds the status of the eventbus resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.Status"),
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Config holds the fininalized configuration of EventBus",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.BusConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.Status", "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.BusConfig"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_NATSBus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NATSBus holds the NATS eventbus information",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"native": {
						SchemaProps: spec.SchemaProps{
							Description: "Native means to bring up a native NATS service",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NativeStrategy"),
						},
					},
					"exotic": {
						SchemaProps: spec.SchemaProps{
							Description: "Exotic holds an exotic NATS config",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NATSConfig", "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.NativeStrategy"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_NATSConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NATSConfig holds the config of NATS",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "NATS host url",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clusterID": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster ID for nats streaming, if it's missing, treat it as NATS server",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"auth": {
						SchemaProps: spec.SchemaProps{
							Description: "Auth strategy, default to AuthStrategyNone",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret for auth",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_NativeStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NativeStrategy indicates to install a native NATS service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Size is the NATS StatefulSet size",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"auth": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"antiAffinity": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"persistence": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.PersistenceStrategy"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1.PersistenceStrategy"},
	}
}

func schema_pkg_apis_eventbus_v1alpha1_PersistenceStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PersistenceStrategy defines the strategy of persistence",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"storageClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessMode": {
						SchemaProps: spec.SchemaProps{
							Description: "Available access modes such as ReadWriteOnce, ReadWriteMany https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Volume size, e.g. 10Gi",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/api/resource.Quantity"},
	}
}
