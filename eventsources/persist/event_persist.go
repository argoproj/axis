package persist

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/argoproj/argo-events/pkg/apis/eventsource/v1alpha1"
)

type EventPersist interface {
	Save(event *Event) error
	Get(key string) (*Event, error)
	IsEnabled() bool
}

type Event struct {
	EventSource  string
	EventPayload string
}

type ConfigMapPersist struct {
	kubeClient kubernetes.Interface
	name       string
	namespace  string
	configMap  *v1.ConfigMap
}

func NewConfigMapPersist(client kubernetes.Interface, configmap *v1alpha1.ConfigMapPersistence, namespace string) (EventPersist, error) {
	if configmap == nil {
		return nil, fmt.Errorf("persistence configuration is nil")
	}
	cm, err := client.CoreV1().ConfigMaps(namespace).Get(configmap.Name, metav1.GetOptions{})
	if err != nil {
		if apierr.IsNotFound(err) && configmap.CreateIfNotExist {
			cm = &v1.ConfigMap{}
			cm.Name = configmap.Name
			cm.Namespace = namespace
			cm, err = client.CoreV1().ConfigMaps(namespace).Create(cm)
		}
		return nil, err
	}

	cmp := ConfigMapPersist{
		kubeClient: client,
		name:       configmap.Name,
		namespace:  namespace,
		configMap:  cm,
	}
	return &cmp, nil
}

func (cmp *ConfigMapPersist) IsEnabled() bool {
	return true
}

func (cmp *ConfigMapPersist) Save(event *Event) error {
	if event == nil {
		return fmt.Errorf("event object is nil")
	}
	cm, err := cmp.kubeClient.CoreV1().ConfigMaps(cmp.namespace).Get(cmp.name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if len(cm.Data) == 0 {
		cm.Data = make(map[string]string)
	}
	cm.Data[event.EventSource] = event.EventPayload
	cmp.configMap, err = cmp.kubeClient.CoreV1().ConfigMaps(cmp.namespace).Update(cm)
	if err != nil {
		return err
	}
	return nil
}

func (cmp *ConfigMapPersist) Get(key string) (*Event, error) {
	cm, err := cmp.kubeClient.CoreV1().ConfigMaps(cmp.namespace).Get(cmp.name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	payload, exist := cm.Data[key]
	if !exist {
		return nil, nil
	}
	return &Event{EventSource: key, EventPayload: payload}, nil
}

type NullPersistence struct {
}

func (n *NullPersistence) Save(event *Event) error {
	return nil
}

func (n *NullPersistence) Get(key string) (*Event, error) {
	return nil, nil
}

func (cmp *NullPersistence) IsEnabled() bool {
	return false
}
