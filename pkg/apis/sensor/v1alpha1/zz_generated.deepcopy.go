// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMQP) DeepCopyInto(out *AMQP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMQP.
func (in *AMQP) DeepCopy() *AMQP {
	if in == nil {
		return nil
	}
	out := new(AMQP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ARN) DeepCopyInto(out *ARN) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ARN.
func (in *ARN) DeepCopy() *ARN {
	if in == nil {
		return nil
	}
	out := new(ARN)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactLocation) DeepCopyInto(out *ArtifactLocation) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		if *in == nil {
			*out = nil
		} else {
			*out = new(S3Artifact)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactLocation.
func (in *ArtifactLocation) DeepCopy() *ArtifactLocation {
	if in == nil {
		return nil
	}
	out := new(ArtifactLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactSignal) DeepCopyInto(out *ArtifactSignal) {
	*out = *in
	in.ArtifactLocation.DeepCopyInto(&out.ArtifactLocation)
	in.NotificationStream.DeepCopyInto(&out.NotificationStream)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactSignal.
func (in *ArtifactSignal) DeepCopy() *ArtifactSignal {
	if in == nil {
		return nil
	}
	out := new(ArtifactSignal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalendarSignal) DeepCopyInto(out *CalendarSignal) {
	*out = *in
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalendarSignal.
func (in *CalendarSignal) DeepCopy() *CalendarSignal {
	if in == nil {
		return nil
	}
	out := new(CalendarSignal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EscalationPolicy) DeepCopyInto(out *EscalationPolicy) {
	*out = *in
	in.Message.DeepCopyInto(&out.Message)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EscalationPolicy.
func (in *EscalationPolicy) DeepCopy() *EscalationPolicy {
	if in == nil {
		return nil
	}
	out := new(EscalationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionKind) DeepCopyInto(out *GroupVersionKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionKind.
func (in *GroupVersionKind) DeepCopy() *GroupVersionKind {
	if in == nil {
		return nil
	}
	out := new(GroupVersionKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kafka) DeepCopyInto(out *Kafka) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kafka.
func (in *Kafka) DeepCopy() *Kafka {
	if in == nil {
		return nil
	}
	out := new(Kafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTT) DeepCopyInto(out *MQTT) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTT.
func (in *MQTT) DeepCopy() *MQTT {
	if in == nil {
		return nil
	}
	out := new(MQTT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Message) DeepCopyInto(out *Message) {
	*out = *in
	in.Stream.DeepCopyInto(&out.Stream)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Message.
func (in *Message) DeepCopy() *Message {
	if in == nil {
		return nil
	}
	out := new(Message)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATS) DeepCopyInto(out *NATS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATS.
func (in *NATS) DeepCopy() *NATS {
	if in == nil {
		return nil
	}
	out := new(NATS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.ResolvedAt.DeepCopyInto(&out.ResolvedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFilter) DeepCopyInto(out *ResourceFilter) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.CreatedBy.DeepCopyInto(&out.CreatedBy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFilter.
func (in *ResourceFilter) DeepCopy() *ResourceFilter {
	if in == nil {
		return nil
	}
	out := new(ResourceFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceObject) DeepCopyInto(out *ResourceObject) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	if in.ArtifactLocation != nil {
		in, out := &in.ArtifactLocation, &out.ArtifactLocation
		if *in == nil {
			*out = nil
		} else {
			*out = new(ArtifactLocation)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceObject.
func (in *ResourceObject) DeepCopy() *ResourceObject {
	if in == nil {
		return nil
	}
	out := new(ResourceObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSignal) DeepCopyInto(out *ResourceSignal) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		if *in == nil {
			*out = nil
		} else {
			*out = new(ResourceFilter)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSignal.
func (in *ResourceSignal) DeepCopy() *ResourceSignal {
	if in == nil {
		return nil
	}
	out := new(ResourceSignal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryStrategy) DeepCopyInto(out *RetryStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryStrategy.
func (in *RetryStrategy) DeepCopy() *RetryStrategy {
	if in == nil {
		return nil
	}
	out := new(RetryStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Artifact) DeepCopyInto(out *S3Artifact) {
	*out = *in
	in.S3Bucket.DeepCopyInto(&out.S3Bucket)
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		if *in == nil {
			*out = nil
		} else {
			*out = new(ARN)
			**out = **in
		}
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		if *in == nil {
			*out = nil
		} else {
			*out = new(S3Filter)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Artifact.
func (in *S3Artifact) DeepCopy() *S3Artifact {
	if in == nil {
		return nil
	}
	out := new(S3Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	in.AccessKey.DeepCopyInto(&out.AccessKey)
	in.SecretKey.DeepCopyInto(&out.SecretKey)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Filter) DeepCopyInto(out *S3Filter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Filter.
func (in *S3Filter) DeepCopy() *S3Filter {
	if in == nil {
		return nil
	}
	out := new(S3Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sensor) DeepCopyInto(out *Sensor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sensor.
func (in *Sensor) DeepCopy() *Sensor {
	if in == nil {
		return nil
	}
	out := new(Sensor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sensor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorList) DeepCopyInto(out *SensorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sensor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorList.
func (in *SensorList) DeepCopy() *SensorList {
	if in == nil {
		return nil
	}
	out := new(SensorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SensorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorSpec) DeepCopyInto(out *SensorSpec) {
	*out = *in
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]Signal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Escalation.DeepCopyInto(&out.Escalation)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorSpec.
func (in *SensorSpec) DeepCopy() *SensorSpec {
	if in == nil {
		return nil
	}
	out := new(SensorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorStatus) DeepCopyInto(out *SensorStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.ResolvedAt.DeepCopyInto(&out.ResolvedAt)
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]NodeStatus, len(*in))
		for key, val := range *in {
			newVal := new(NodeStatus)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorStatus.
func (in *SensorStatus) DeepCopy() *SensorStatus {
	if in == nil {
		return nil
	}
	out := new(SensorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signal) DeepCopyInto(out *Signal) {
	*out = *in
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		if *in == nil {
			*out = nil
		} else {
			*out = new(NATS)
			**out = **in
		}
	}
	if in.MQTT != nil {
		in, out := &in.MQTT, &out.MQTT
		if *in == nil {
			*out = nil
		} else {
			*out = new(MQTT)
			**out = **in
		}
	}
	if in.AMQP != nil {
		in, out := &in.AMQP, &out.AMQP
		if *in == nil {
			*out = nil
		} else {
			*out = new(AMQP)
			**out = **in
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		if *in == nil {
			*out = nil
		} else {
			*out = new(Kafka)
			**out = **in
		}
	}
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		if *in == nil {
			*out = nil
		} else {
			*out = new(ArtifactSignal)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Calendar != nil {
		in, out := &in.Calendar, &out.Calendar
		if *in == nil {
			*out = nil
		} else {
			*out = new(CalendarSignal)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		if *in == nil {
			*out = nil
		} else {
			*out = new(ResourceSignal)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebhookSignal)
			**out = **in
		}
	}
	in.Constraints.DeepCopyInto(&out.Constraints)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signal.
func (in *Signal) DeepCopy() *Signal {
	if in == nil {
		return nil
	}
	out := new(Signal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignalConstraints) DeepCopyInto(out *SignalConstraints) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignalConstraints.
func (in *SignalConstraints) DeepCopy() *SignalConstraints {
	if in == nil {
		return nil
	}
	out := new(SignalConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		if *in == nil {
			*out = nil
		} else {
			*out = new(NATS)
			**out = **in
		}
	}
	if in.MQTT != nil {
		in, out := &in.MQTT, &out.MQTT
		if *in == nil {
			*out = nil
		} else {
			*out = new(MQTT)
			**out = **in
		}
	}
	if in.AMQP != nil {
		in, out := &in.AMQP, &out.AMQP
		if *in == nil {
			*out = nil
		} else {
			*out = new(AMQP)
			**out = **in
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		if *in == nil {
			*out = nil
		} else {
			*out = new(Kafka)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeConstraints) DeepCopyInto(out *TimeConstraints) {
	*out = *in
	in.Start.DeepCopyInto(&out.Start)
	in.Stop.DeepCopyInto(&out.Stop)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeConstraints.
func (in *TimeConstraints) DeepCopy() *TimeConstraints {
	if in == nil {
		return nil
	}
	out := new(TimeConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		if *in == nil {
			*out = nil
		} else {
			*out = new(ResourceObject)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		if *in == nil {
			*out = nil
		} else {
			*out = new(Message)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RetryStrategy != nil {
		in, out := &in.RetryStrategy, &out.RetryStrategy
		if *in == nil {
			*out = nil
		} else {
			*out = new(RetryStrategy)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSignal) DeepCopyInto(out *WebhookSignal) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSignal.
func (in *WebhookSignal) DeepCopy() *WebhookSignal {
	if in == nil {
		return nil
	}
	out := new(WebhookSignal)
	in.DeepCopyInto(out)
	return out
}
