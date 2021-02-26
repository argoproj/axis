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

package common

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/pkg/apis/common.Amount":     schema_argo_events_pkg_apis_common_Amount(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.Backoff":    schema_argo_events_pkg_apis_common_Backoff(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.BasicAuth":  schema_argo_events_pkg_apis_common_BasicAuth(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.Condition":  schema_argo_events_pkg_apis_common_Condition(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.Metadata":   schema_argo_events_pkg_apis_common_Metadata(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.Resource":   schema_argo_events_pkg_apis_common_Resource(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.S3Artifact": schema_argo_events_pkg_apis_common_S3Artifact(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.S3Bucket":   schema_argo_events_pkg_apis_common_S3Bucket(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.S3Filter":   schema_argo_events_pkg_apis_common_S3Filter(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.Status":     schema_argo_events_pkg_apis_common_Status(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.TLSConfig":  schema_argo_events_pkg_apis_common_TLSConfig(ref),
	}
}

func schema_argo_events_pkg_apis_common_Amount(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Amount represent a numeric amount.",
				Type:        Amount{}.OpenAPISchemaType(),
				Format:      Amount{}.OpenAPISchemaFormat(),
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_Backoff(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Backoff for an operation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"duration": {
						SchemaProps: spec.SchemaProps{
							Description: "The initial duration such as \"1s\", \"3m\"",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"factor": {
						SchemaProps: spec.SchemaProps{
							Description: "Duration is multiplied by factor each iteration",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/common.Amount"),
						},
					},
					"jitter": {
						SchemaProps: spec.SchemaProps{
							Description: "The amount of jitter applied each iteration",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/common.Amount"),
						},
					},
					"steps": {
						SchemaProps: spec.SchemaProps{
							Description: "Exit with error after this many steps",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.Amount"},
	}
}

func schema_argo_events_pkg_apis_common_BasicAuth(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BasicAuth contains the reference to K8s secrets that holds the username and password",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"username": {
						SchemaProps: spec.SchemaProps{
							Description: "Username refers to the Kubernetes secret that holds the username required for basic auth.",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"password": {
						SchemaProps: spec.SchemaProps{
							Description: "Password refers to the Kubernetes secret that holds the password required for basic auth.",
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

func schema_argo_events_pkg_apis_common_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Condition contains details about resource state",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Condition type.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Condition status, True, False or Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "Unique, this should be a short, machine understandable string that gives the reason for condition's last transition. For example, \"ImageNotFound\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human-readable message indicating details about last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_argo_events_pkg_apis_common_Metadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Metadata holds the annotations and labels of an event source pod",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_Resource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Resource represent arbitrary structured data.",
				Type:        Resource{}.OpenAPISchemaType(),
				Format:      Resource{}.OpenAPISchemaFormat(),
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_S3Artifact(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Artifact contains information about an S3 connection and bucket",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.S3Bucket"),
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"insecure": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"accessKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"secretKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"events": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"filter": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.S3Filter"),
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"endpoint", "bucket", "accessKey", "secretKey"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.S3Bucket", "github.com/argoproj/argo-events/pkg/apis/common.S3Filter", "k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_argo_events_pkg_apis_common_S3Bucket(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Bucket contains information to describe an S3 Bucket",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"key": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_S3Filter(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Filter represents filters to apply to bucket notifications for specifying constraints on objects",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"suffix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"prefix", "suffix"},
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_Status(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Status is a common structure which can be used for Status field.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions are the latest available observations of a resource's current state.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.Condition"},
	}
}

func schema_argo_events_pkg_apis_common_TLSConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TLSConfig refers to TLS configuration for a client.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"caCertSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "CACertSecret refers to the secret that contains the CA cert",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"clientCertSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "ClientCertSecret refers to the secret that contains the client cert",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"clientKeySecret": {
						SchemaProps: spec.SchemaProps{
							Description: "ClientKeySecret refers to the secret that contains the client key",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"caCertPath": {
						SchemaProps: spec.SchemaProps{
							Description: "DeprecatedCACertPath refers the file path that contains the CA cert. Deprecated: use CACertSecret instead",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clientCertPath": {
						SchemaProps: spec.SchemaProps{
							Description: "DeprecatedClientCertPath refers the file path that contains client cert. Deprecated: use ClientCertSecret instead",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clientKeyPath": {
						SchemaProps: spec.SchemaProps{
							Description: "DeprecatedClientKeyPath refers the file path that contains client key. Deprecated: use ClientKeySecret instead",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretKeySelector"},
	}
}
