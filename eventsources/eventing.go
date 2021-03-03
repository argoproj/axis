package eventsources

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/common/logging"
	"github.com/argoproj/argo-events/eventbus"
	eventbusdriver "github.com/argoproj/argo-events/eventbus/driver"
	"github.com/argoproj/argo-events/eventsources/sources/amqp"
	"github.com/argoproj/argo-events/eventsources/sources/awssns"
	"github.com/argoproj/argo-events/eventsources/sources/awssqs"
	"github.com/argoproj/argo-events/eventsources/sources/azureeventshub"
	"github.com/argoproj/argo-events/eventsources/sources/calendar"
	"github.com/argoproj/argo-events/eventsources/sources/emitter"
	"github.com/argoproj/argo-events/eventsources/sources/file"
	"github.com/argoproj/argo-events/eventsources/sources/gcppubsub"
	"github.com/argoproj/argo-events/eventsources/sources/generic"
	"github.com/argoproj/argo-events/eventsources/sources/github"
	"github.com/argoproj/argo-events/eventsources/sources/gitlab"
	"github.com/argoproj/argo-events/eventsources/sources/hdfs"
	"github.com/argoproj/argo-events/eventsources/sources/kafka"
	"github.com/argoproj/argo-events/eventsources/sources/minio"
	"github.com/argoproj/argo-events/eventsources/sources/mqtt"
	"github.com/argoproj/argo-events/eventsources/sources/nats"
	"github.com/argoproj/argo-events/eventsources/sources/nsq"
	"github.com/argoproj/argo-events/eventsources/sources/pulsar"
	"github.com/argoproj/argo-events/eventsources/sources/redis"
	"github.com/argoproj/argo-events/eventsources/sources/resource"
	"github.com/argoproj/argo-events/eventsources/sources/slack"
	"github.com/argoproj/argo-events/eventsources/sources/storagegrid"
	"github.com/argoproj/argo-events/eventsources/sources/stripe"
	"github.com/argoproj/argo-events/eventsources/sources/webhook"
	eventsourcemetrics "github.com/argoproj/argo-events/metrics"
	apicommon "github.com/argoproj/argo-events/pkg/apis/common"
	eventbusv1alpha1 "github.com/argoproj/argo-events/pkg/apis/eventbus/v1alpha1"
	"github.com/argoproj/argo-events/pkg/apis/eventsource/v1alpha1"
)

// EventingServer is the server API for Eventing service.
type EventingServer interface {

	// ValidateEventSource validates an event source.
	ValidateEventSource(context.Context) error

	GetEventSourceName() string

	GetEventName() string

	GetEventSourceType() apicommon.EventSourceType

	// Function to start listening events.
	StartListening(ctx context.Context, dispatch func([]byte) error) error
}

// GetEventingServers returns the mapping of event source type and list of eventing servers
func GetEventingServers(eventSource *v1alpha1.EventSource) map[apicommon.EventSourceType][]EventingServer {
	result := make(map[apicommon.EventSourceType][]EventingServer)
	if len(eventSource.Spec.AMQP) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.AMQP {
			servers = append(servers, &amqp.EventListener{EventSourceName: eventSource.Name, EventName: k, AMQPEventSource: v})
		}
		result[apicommon.AMQPEvent] = servers
	}
	if len(eventSource.Spec.AzureEventsHub) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.AzureEventsHub {
			servers = append(servers, &azureeventshub.EventListener{EventSourceName: eventSource.Name, EventName: k, AzureEventsHubEventSource: v})
		}
		result[apicommon.AzureEventsHub] = servers
	}
	if len(eventSource.Spec.Calendar) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Calendar {
			servers = append(servers, &calendar.EventListener{EventSourceName: eventSource.Name, EventName: k, CalendarEventSource: v, Namespace: eventSource.Namespace})
		}
		result[apicommon.CalendarEvent] = servers
	}
	if len(eventSource.Spec.Emitter) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Emitter {
			servers = append(servers, &emitter.EventListener{EventSourceName: eventSource.Name, EventName: k, EmitterEventSource: v})
		}
		result[apicommon.EmitterEvent] = servers
	}
	if len(eventSource.Spec.File) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.File {
			servers = append(servers, &file.EventListener{EventSourceName: eventSource.Name, EventName: k, FileEventSource: v})
		}
		result[apicommon.FileEvent] = servers
	}
	if len(eventSource.Spec.Github) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Github {
			servers = append(servers, &github.EventListener{EventSourceName: eventSource.Name, EventName: k, GithubEventSource: v})
		}
		result[apicommon.GithubEvent] = servers
	}
	if len(eventSource.Spec.Gitlab) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Gitlab {
			servers = append(servers, &gitlab.EventListener{EventSourceName: eventSource.Name, EventName: k, GitlabEventSource: v})
		}
		result[apicommon.GitlabEvent] = servers
	}
	if len(eventSource.Spec.HDFS) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.HDFS {
			servers = append(servers, &hdfs.EventListener{EventSourceName: eventSource.Name, EventName: k, HDFSEventSource: v})
		}
		result[apicommon.HDFSEvent] = servers
	}
	if len(eventSource.Spec.Kafka) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Kafka {
			servers = append(servers, &kafka.EventListener{EventSourceName: eventSource.Name, EventName: k, KafkaEventSource: v})
		}
		result[apicommon.KafkaEvent] = servers
	}
	if len(eventSource.Spec.MQTT) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.MQTT {
			servers = append(servers, &mqtt.EventListener{EventSourceName: eventSource.Name, EventName: k, MQTTEventSource: v})
		}
		result[apicommon.MQTTEvent] = servers
	}
	if len(eventSource.Spec.Minio) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Minio {
			servers = append(servers, &minio.EventListener{EventSourceName: eventSource.Name, EventName: k, MinioEventSource: v})
		}
		result[apicommon.MinioEvent] = servers
	}
	if len(eventSource.Spec.NATS) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.NATS {
			servers = append(servers, &nats.EventListener{EventSourceName: eventSource.Name, EventName: k, NATSEventSource: v})
		}
		result[apicommon.NATSEvent] = servers
	}
	if len(eventSource.Spec.NSQ) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.NSQ {
			servers = append(servers, &nsq.EventListener{EventSourceName: eventSource.Name, EventName: k, NSQEventSource: v})
		}
		result[apicommon.NSQEvent] = servers
	}
	if len(eventSource.Spec.PubSub) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.PubSub {
			servers = append(servers, &gcppubsub.EventListener{EventSourceName: eventSource.Name, EventName: k, PubSubEventSource: v})
		}
		result[apicommon.PubSubEvent] = servers
	}
	if len(eventSource.Spec.Redis) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Redis {
			servers = append(servers, &redis.EventListener{EventSourceName: eventSource.Name, EventName: k, RedisEventSource: v})
		}
		result[apicommon.RedisEvent] = servers
	}
	if len(eventSource.Spec.SNS) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.SNS {
			servers = append(servers, &awssns.EventListener{EventSourceName: eventSource.Name, EventName: k, SNSEventSource: v})
		}
		result[apicommon.SNSEvent] = servers
	}
	if len(eventSource.Spec.SQS) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.SQS {
			servers = append(servers, &awssqs.EventListener{EventSourceName: eventSource.Name, EventName: k, SQSEventSource: v})
		}
		result[apicommon.SQSEvent] = servers
	}
	if len(eventSource.Spec.Slack) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Slack {
			servers = append(servers, &slack.EventListener{EventSourceName: eventSource.Name, EventName: k, SlackEventSource: v})
		}
		result[apicommon.SlackEvent] = servers
	}
	if len(eventSource.Spec.StorageGrid) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.StorageGrid {
			servers = append(servers, &storagegrid.EventListener{EventSourceName: eventSource.Name, EventName: k, StorageGridEventSource: v})
		}
		result[apicommon.StorageGridEvent] = servers
	}
	if len(eventSource.Spec.Stripe) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Stripe {
			servers = append(servers, &stripe.EventListener{EventSourceName: eventSource.Name, EventName: k, StripeEventSource: v})
		}
		result[apicommon.StripeEvent] = servers
	}
	if len(eventSource.Spec.Webhook) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Webhook {
			servers = append(servers, &webhook.EventListener{EventSourceName: eventSource.Name, EventName: k, WebhookContext: v})
		}
		result[apicommon.WebhookEvent] = servers
	}
	if len(eventSource.Spec.Resource) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Resource {
			servers = append(servers, &resource.EventListener{EventSourceName: eventSource.Name, EventName: k, ResourceEventSource: v})
		}
		result[apicommon.ResourceEvent] = servers
	}
	if len(eventSource.Spec.Pulsar) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Pulsar {
			servers = append(servers, &pulsar.EventListener{EventSourceName: eventSource.Name, EventName: k, PulsarEventSource: v})
		}
		result[apicommon.PulsarEvent] = servers
	}
	if len(eventSource.Spec.Generic) != 0 {
		servers := []EventingServer{}
		for k, v := range eventSource.Spec.Generic {
			servers = append(servers, &generic.EventListener{EventSourceName: eventSource.Name, EventName: k, GenericEventSource: v})
		}
		result[apicommon.GenericEvent] = servers
	}
	return result
}

// EventSourceAdaptor is the adaptor for eventsource service
type EventSourceAdaptor struct {
	eventSource     *v1alpha1.EventSource
	eventBusConfig  *eventbusv1alpha1.BusConfig
	eventBusSubject string
	hostname        string

	eventBusConn eventbusdriver.Connection

	metrics *eventsourcemetrics.Metrics
}

// NewEventSourceAdaptor returns a new EventSourceAdaptor
func NewEventSourceAdaptor(eventSource *v1alpha1.EventSource, eventBusConfig *eventbusv1alpha1.BusConfig, eventBusSubject, hostname string, metrics *eventsourcemetrics.Metrics) *EventSourceAdaptor {
	return &EventSourceAdaptor{
		eventSource:     eventSource,
		eventBusConfig:  eventBusConfig,
		eventBusSubject: eventBusSubject,
		hostname:        hostname,
		metrics:         metrics,
	}
}

// Start function
func (e *EventSourceAdaptor) Start(ctx context.Context) error {
	logger := logging.FromContext(ctx).Desugar()
	logger.Info("Starting event source server...")
	servers := GetEventingServers(e.eventSource)
	clientID := generateClientID(e.hostname)
	driver, err := eventbus.GetDriver(ctx, *e.eventBusConfig, e.eventBusSubject, clientID)
	if err != nil {
		logger.Error("failed to get eventbus driver", zap.Error(err))
		return err
	}
	if err = common.Connect(&common.DefaultBackoff, func() error {
		e.eventBusConn, err = driver.Connect()
		return err
	}); err != nil {
		logger.Error("failed to connect to eventbus", zap.Error(err))
		return err
	}
	defer e.eventBusConn.Close()

	cctx, cancel := context.WithCancel(ctx)
	connWG := &sync.WaitGroup{}

	// Daemon to reconnect
	connWG.Add(1)
	go func() {
		defer connWG.Done()
		logger.Info("starting eventbus connection daemon...")
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-cctx.Done():
				logger.Info("exiting eventbus connection daemon...")
				return
			case <-ticker.C:
				if e.eventBusConn == nil || e.eventBusConn.IsClosed() {
					logger.Info("NATS connection lost, reconnecting...")
					// Regenerate the client ID to avoid the issue that NAT server still thinks the client is alive.
					clientID := generateClientID(e.hostname)
					driver, err := eventbus.GetDriver(cctx, *e.eventBusConfig, e.eventBusSubject, clientID)
					if err != nil {
						logger.Error("failed to get eventbus driver during reconnection", zap.Error(err))
						continue
					}
					e.eventBusConn, err = driver.Connect()
					if err != nil {
						logger.Error("failed to reconnect to eventbus", zap.Error(err))
						continue
					}
					logger.Info("reconnected to eventbus successfully")
				}
			}
		}
	}()

	wg := &sync.WaitGroup{}
	for _, ss := range servers {
		for _, server := range ss {
			// Validation has been done in eventsource-controller, it's harmless to do it again here.
			err := server.ValidateEventSource(cctx)
			if err != nil {
				logger.Error("Validation failed", zap.Error(err), zap.Any(logging.LabelEventName,
					server.GetEventName()), zap.Any(logging.LabelEventSourceType, server.GetEventSourceType()))
				// Continue starting other event services instead of failing all of them
				continue
			}
			wg.Add(1)
			go func(s EventingServer) {
				e.metrics.IncRunningServices(s.GetEventSourceName())
				defer e.metrics.DecRunningServices(s.GetEventSourceName())
				defer wg.Done()
				duration := apicommon.FromString("1s")
				factor := apicommon.NewAmount("1")
				jitter := apicommon.NewAmount("30")
				backoff := apicommon.Backoff{
					Steps:    10,
					Duration: &duration,
					Factor:   &factor,
					Jitter:   &jitter,
				}
				if err = common.Connect(&backoff, func() error {
					return s.StartListening(cctx, func(data []byte) error {
						event := cloudevents.NewEvent()
						event.SetID(fmt.Sprintf("%x", uuid.New()))
						event.SetType(string(s.GetEventSourceType()))
						event.SetSource(s.GetEventSourceName())
						event.SetSubject(s.GetEventName())
						event.SetTime(time.Now())
						err := event.SetData(cloudevents.ApplicationJSON, data)
						if err != nil {
							return err
						}
						eventBody, err := json.Marshal(event)
						if err != nil {
							return err
						}
						if e.eventBusConn == nil || e.eventBusConn.IsClosed() {
							return errors.New("failed to publish event, eventbus connection closed")
						}
						if err = driver.Publish(e.eventBusConn, eventBody); err != nil {
							logger.Error("failed to publish an event", zap.Error(err), zap.String(logging.LabelEventName,
								s.GetEventName()), zap.Any(logging.LabelEventSourceType, s.GetEventSourceType()))
							e.metrics.EventSentFailed(s.GetEventSourceName(), s.GetEventName())
							return err
						}
						logger.Info("succeeded to publish an event", zap.Error(err), zap.String(logging.LabelEventName,
							s.GetEventName()), zap.Any(logging.LabelEventSourceType, s.GetEventSourceType()), zap.String("eventID", event.ID()))
						e.metrics.EventSent(s.GetEventSourceName(), s.GetEventName())
						return nil
					})
				}); err != nil {
					logger.Error("failed to start listening eventsource", zap.Any(logging.LabelEventSourceType,
						s.GetEventSourceType()), zap.Any(logging.LabelEventName, s.GetEventName()), zap.Error(err))
				}
			}(server)
		}
	}
	logger.Info("Eventing server started.")

	eventServersWGDone := make(chan bool)
	go func() {
		wg.Wait()
		close(eventServersWGDone)
	}()

	for {
		select {
		case <-ctx.Done():
			logger.Info("Shutting down...")
			cancel()
			<-eventServersWGDone
			connWG.Wait()
			return nil
		case <-eventServersWGDone:
			logger.Error("Erroring out, no active event server running")
			cancel()
			connWG.Wait()
			return errors.New("no active event server running")
		}
	}
}

func generateClientID(hostname string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	clientID := fmt.Sprintf("client-%s-%v", strings.ReplaceAll(hostname, ".", "_"), r1.Intn(1000))
	return clientID
}
