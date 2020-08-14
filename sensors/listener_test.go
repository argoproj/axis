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

package sensors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
)

var (
	fakeTrigger = &v1alpha1.Trigger{
		Template: &v1alpha1.TriggerTemplate{
			Name: "fake-trigger",
			K8s: &v1alpha1.StandardK8STrigger{
				GroupVersionResource: metav1.GroupVersionResource{
					Group:    "apps",
					Version:  "v1",
					Resource: "deployments",
				},
			},
		},
	}

	sensorObj = &v1alpha1.Sensor{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-sensor",
			Namespace: "fake",
		},
		Spec: v1alpha1.SensorSpec{
			Triggers: []v1alpha1.Trigger{
				*fakeTrigger,
			},
		},
	}
)

func TestGetDependencyExpression(t *testing.T) {
	t.Run("get simple expression", func(t *testing.T) {
		obj := sensorObj.DeepCopy()
		obj.Spec.Dependencies = []v1alpha1.EventDependency{
			{
				Name:            "dep1",
				EventSourceName: "webhook",
				EventName:       "example-1",
			},
		}
		sensorCtx := &SensorContext{
			Sensor: obj,
		}
		expr, err := sensorCtx.getDependencyExpression(context.Background(), *fakeTrigger)
		assert.NoError(t, err)
		assert.Equal(t, "dep1", expr)
	})

	t.Run("get two deps expression", func(t *testing.T) {
		obj := sensorObj.DeepCopy()
		obj.Spec.Dependencies = []v1alpha1.EventDependency{
			{
				Name:            "dep1",
				EventSourceName: "webhook",
				EventName:       "example-1",
			},
			{
				Name:            "dep2",
				EventSourceName: "webhook2",
				EventName:       "example-2",
			},
		}
		sensorCtx := &SensorContext{
			Sensor: obj,
		}
		expr, err := sensorCtx.getDependencyExpression(context.Background(), *fakeTrigger)
		assert.NoError(t, err)
		assert.Equal(t, "dep1 && dep2", expr)
	})

	t.Run("get complex expression", func(t *testing.T) {
		obj := sensorObj.DeepCopy()
		obj.Spec.Dependencies = []v1alpha1.EventDependency{
			{
				Name:            "dep1",
				EventSourceName: "webhook",
				EventName:       "example-1",
			},
			{
				Name:            "dep1a",
				EventSourceName: "webhook",
				EventName:       "example-1a",
			},
			{
				Name:            "dep2",
				EventSourceName: "webhook2",
				EventName:       "example-2",
			},
		}
		sensorCtx := &SensorContext{
			Sensor: obj,
		}
		obj.Spec.DependencyGroups = []v1alpha1.DependencyGroup{
			{Name: "group-1", Dependencies: []string{"dep1", "dep1a"}},
			{Name: "group-2", Dependencies: []string{"dep2"}},
		}
		obj.Spec.DeprecatedCircuit = "((group-2) || group-1)"
		trig := fakeTrigger.DeepCopy()
		trig.Template.DeprecatedSwitch = &v1alpha1.TriggerSwitch{
			Any: []string{"group-1"},
		}
		expr, err := sensorCtx.getDependencyExpression(context.Background(), *trig)
		assert.NoError(t, err)
		assert.Equal(t, "dep1 && dep1a", expr)
	})

	t.Run("get conditions expression", func(t *testing.T) {
		obj := sensorObj.DeepCopy()
		obj.Spec.Dependencies = []v1alpha1.EventDependency{
			{
				Name:            "dep-1",
				EventSourceName: "webhook",
				EventName:       "example-1",
			},
			{
				Name:            "dep_1a",
				EventSourceName: "webhook",
				EventName:       "example-1a",
			},
			{
				Name:            "dep-2",
				EventSourceName: "webhook2",
				EventName:       "example-2",
			},
			{
				Name:            "dep-3",
				EventSourceName: "webhook3",
				EventName:       "example-3",
			},
			{
				Name:            "dep-4",
				EventSourceName: "webhook4",
				EventName:       "example-4",
			},
			{
				Name:            "dep_5",
				EventSourceName: "webhook5",
				EventName:       "example-5",
			},
			{
				Name:            "dep-6",
				EventSourceName: "webhook6",
				EventName:       "example-6",
			},
		}
		sensorCtx := &SensorContext{
			Sensor: obj,
		}
		obj.Spec.DependencyGroups = []v1alpha1.DependencyGroup{
			{Name: "group-1", Dependencies: []string{"dep-1", "dep_1a"}},
			{Name: "group_2", Dependencies: []string{"dep-2"}},
		}
		obj.Spec.DependencyAliases = map[string]string{
			"alias-01": "(dep-4 &&dep_5) ||dep-6",
		}
		trig := fakeTrigger.DeepCopy()
		trig.Template.Conditions = "group-1 || group_2 || dep-3 || alias-01"
		expr, err := sensorCtx.getDependencyExpression(context.Background(), *trig)
		assert.NoError(t, err)
		assert.Equal(t, "dep-6 || dep-3 || dep-2 || (dep-4 && dep_5) || (dep-1 && dep_1a)", expr)
	})
}
