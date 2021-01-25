package validator

import (
	"context"

	"go.uber.org/zap"
	admissionv1 "k8s.io/api/admission/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/common/logging"
	sensorcontroller "github.com/argoproj/argo-events/controllers/sensor"
	sensorv1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	eventbusclient "github.com/argoproj/argo-events/pkg/client/eventbus/clientset/versioned"
	eventsourceclient "github.com/argoproj/argo-events/pkg/client/eventsource/clientset/versioned"
	sensorclient "github.com/argoproj/argo-events/pkg/client/sensor/clientset/versioned"
)

type sensor struct {
	client            kubernetes.Interface
	eventBusClient    eventbusclient.Interface
	eventSourceClient eventsourceclient.Interface
	sensorClient      sensorclient.Interface

	oldSensor *sensorv1alpha1.Sensor
	newSensor *sensorv1alpha1.Sensor
}

// NewSensorValidator returns a validator for Sensor
func NewSensorValidator(client kubernetes.Interface, ebClient eventbusclient.Interface,
	esClient eventsourceclient.Interface, sClient sensorclient.Interface, old, new *sensorv1alpha1.Sensor) Validator {
	return &sensor{client: client, eventBusClient: ebClient, eventSourceClient: esClient, sensorClient: sClient, oldSensor: old, newSensor: new}
}

func (s *sensor) ValidateCreate(ctx context.Context) *admissionv1.AdmissionResponse {
	log := logging.FromContext(ctx)
	if err := sensorcontroller.ValidateSensor(s.newSensor); err != nil {
		return DeniedResponse(err.Error())
	}
	ebName := s.newSensor.Spec.EventBusName
	if ebName == "" {
		ebName = common.DefaultEventBusName
	}
	eb, err := s.eventBusClient.ArgoprojV1alpha1().EventBus(s.newSensor.Namespace).Get(ctx, ebName, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return DeniedResponse("An EventBus named \"%s\" needs to be created first", ebName)
		}
		log.Errorw("failed to retrieve EventBus", zap.Error(err))
		return DeniedResponse("Failed to retrieve the EventBus, %s", err.Error())
	}
	if !eb.Status.IsReady() {
		return DeniedResponse("EventBus \"%s\" is not in a good shape", ebName)
	}
	return AllowedResponse()
}

func (s *sensor) ValidateUpdate(ctx context.Context) *admissionv1.AdmissionResponse {
	if s.oldSensor.Generation == s.newSensor.Generation {
		AllowedResponse()
	}
	return s.ValidateCreate(ctx)
}

func (s *sensor) ValidateDelete(ctx context.Context) *admissionv1.AdmissionResponse {
	return AllowedResponse()
}
