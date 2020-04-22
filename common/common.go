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

package common

import (
	"github.com/pkg/errors"
)

// Defaults
const (
	// StandardTimeFormat is time format reference for golang
	StandardTimeFormat = "2006-01-02 15:04:05"
	// StandardYYYYMMDDFormat formats date in yyyy-mm-dd format
	StandardYYYYMMDDFormat = "2006-01-02"
	// DefaultControllerNamespace is the default namespace where the sensor and gateways controllers are installed
	DefaultControllerNamespace = "argo-events"
)

// Environment variables
const (
	// EnvVarKubeConfig is the path to the Kubernetes configuration
	EnvVarKubeConfig = "KUBE_CONFIG"
	// EnvVarDebugLog is the env var to turn on the debug mode for logging
	EnvVarDebugLog = "DEBUG_LOG"
)

// Controller environment variables
const (
	// EnvVarControllerConfigMap contains name of the configmap to retrieve controller configuration from
	EnvVarControllerConfigMap = "CONTROLLER_CONFIG_MAP"
	// EnvVarControllerInstanceID is used to get controller instance id
	EnvVarControllerInstanceID = "CONTROLLER_INSTANCE_ID"
	// EnvVarResourceName refers env var for name of the resource
	EnvVarResourceName = "NAME"
	// EnvVarNamespace refers to a K8s namespace
	EnvVarNamespace = "NAMESPACE"
)

// Controller labels
const (
	// LabelGatewayName is the label for the K8s resource name
	LabelResourceName = "resource-name"
)

const (
	// GatewayControllerConfigMapKey is the key in the configmap to retrieve controller configuration from.
	// Content encoding is expected to be YAML.
	ControllerConfigMapKey = "config"
	// SensorTemplateConfigKey is the key in the configmap to retrieve Sensor deployment template spec.
	SensorTemplateConfigKey = "sensor-template"
)

// Sensor constants
const (
	// SensorName refers env var for name of sensor
	SensorName = "SENSOR_NAME"
	// SensorNamespace is used to get namespace where sensors are deployed
	SensorNamespace = "SENSOR_NAMESPACE"
	// LabelSensorName is label for sensor name
	LabelSensorName = "sensor-name"
	// Port for the sensor server to listen events on
	SensorServerPort = 9300
)

// Gateway constants
const (
	// LabelEventSourceName is the label for a event source in gateway
	LabelEventSourceName    = "event-source-name"
	EnvVarGatewayServerPort = "GATEWAY_SERVER_PORT"
	// ProcessorPort is the default port for the gateway event processor server to run on.
	GatewayProcessorPort = "9300"
)

const (
	// EnvVarEventSource refers to event source name
	EnvVarEventSource = "EVENT_SOURCE"
	// AnnotationResourceSpecHash is the annotation of a K8s resource spec hash
	AnnotationResourceSpecHash = "resource-spec-hash"
)

var (
	ErrNilEventSource = errors.New("event source can't be nil")
)

// Miscellaneous Labels
const (
	// LabelEventSource is label for event name
	LabelEventSource = "event-source"
	// LabelOwnerName is the label for resource owner name
	LabelOwnerName = "owner-name"
	// LabelObjectName is the label for object name
	LabelObjectName = "object-name"
)

// various supported media types
const (
	MediaTypeJSON string = "application/json"
	MediaTypeYAML string = "application/yaml"
)
