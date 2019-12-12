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
	"testing"

	"github.com/argoproj/argo-events/common"
	"github.com/stretchr/testify/assert"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSetObjectMeta(t *testing.T) {
	owner := appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-deployment",
			Namespace: "fake-namespace",
		},
	}
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "fake-pod",
		},
	}

	err := SetObjectMeta(&owner, &pod, owner.GroupVersionKind())
	assert.Nil(t, err)
	assert.Equal(t, "fake-namespace", pod.Namespace)
	assert.Equal(t, owner.GroupVersionKind().Kind, pod.OwnerReferences[0].Kind)
	assert.NotEmpty(t, pod.Annotations[common.AnnotationResourceSpecHash])
	assert.NotEmpty(t, pod.Labels)
	assert.Equal(t, owner.Name, pod.Labels[common.LabelOwnerName])
}
