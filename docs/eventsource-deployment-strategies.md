# EventSource Deployment Strategies

EventSource controller creates a k8s deployment for each EventSource object to
watch the events. Some of the event source types do not allow multiple live
clients with same attributes (i.e. multiple clients with same `clientID`
connecting to a NATS server), or multiple event source PODs will generate
duplicated events to downstream, so the deployment strategy and replica numbers
are different for different event sources.

## Rolling Update Strategy

`Rolling Update` strategy is applied to the following EventSource types:

- AMQP
- AWS SNS
- AWS SQS
- Github
- Gitlab
- NetApp Storage GRID
- Slack
- Stripe
- Webhook

## Recreate Strategy

`Recreate` strategy is applied to the following EventSource types:

- Azure Events Hub
- Kafka
- GCP PubSub
- File
- HDFS
- NATS
- Minio
- MQTT
- Emitter
- NSQ
- Pulsar
- Redis
- Resource
- Calendar

### Replicas

The field `replica` for EventSource of these `Recreate` types are ignored, the
deployment is always created with `replica=1`.

**Please DO NOT manually scale up the replicas, that will bring unexpected
behaviors!**