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

package sensor

import (
	"testing"

	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	"github.com/ghodss/yaml"
	"github.com/smartystreets/goconvey/convey"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

var sensorStr = `
apiVersion: argoproj.io/v1alpha1
kind: Sensor
metadata:
  name: artifact-sensor
  namespace: argo-events
  labels:
    sensors.argoproj.io/sensor-controller-instanceid: argo-events
spec:
  deploySpec:
    containers:
      - name: "sensor"
        image: "argoproj/sensor:v0.7"
        imagePullPolicy: Always
    serviceAccountName: argo-events-sa
  dependencies:
    - name: artifact-gateway:input
  eventProtocol:
    type: "HTTP"
    http:
      port: "9300"
  triggers:
    - name: artifact-workflow-trigger
      resource:
        namespace: argo-events
        group: argoproj.io
        version: v1alpha1
        kind: Workflow
        source:
          inline: |
              apiVersion: argoproj.io/v1alpha1
              kind: Workflow
              metadata:
                generateName: hello-world-
              spec:
                entrypoint: whalesay
                templates:
                  -
                    container:
                      args:
                        - "hello world"
                      command:
                        - cowsay
                      image: "docker/whalesay:latest"
                    name: whalesay
`

func getSensor() (*v1alpha1.Sensor, error) {
	var sensor *v1alpha1.Sensor
	err := yaml.Unmarshal([]byte(sensorStr), &sensor)
	return sensor, err
}

func TestSensorOperations(t *testing.T) {
	done := make(chan struct{})
	convey.Convey("Given a sensor, parse it", t, func() {
		sensor, err := getSensor()
		convey.So(err, convey.ShouldBeNil)
		convey.So(sensor, convey.ShouldNotBeNil)

		controller := getSensorController()
		soc := newSensorOperationCtx(sensor, controller)
		convey.So(soc, convey.ShouldNotBeNil)

		convey.Convey("Create the sensor", func() {
			sensor, err = controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Create(sensor)
			convey.So(err, convey.ShouldBeNil)
			convey.So(sensor, convey.ShouldNotBeNil)

			convey.Convey("Operate on a new sensor", func() {
				soc.markSensorPhase(v1alpha1.NodePhaseNew, false, "test")
				err := soc.operate()
				convey.So(err, convey.ShouldBeNil)

				convey.Convey("Sensor should be marked as active with it's nodes initialized", func() {
					sensor, err = controller.sensorClientset.ArgoprojV1alpha1().Sensors(soc.s.Namespace).Get(soc.s.Name, metav1.GetOptions{})
					convey.So(err, convey.ShouldBeNil)
					convey.So(sensor, convey.ShouldNotBeNil)
					convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)

					for _, node := range soc.s.Status.Nodes {
						switch node.Type {
						case v1alpha1.NodeTypeEventDependency:
							convey.So(node.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
						case v1alpha1.NodeTypeDependencyGroup:
							convey.So(node.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
						case v1alpha1.NodeTypeTrigger:
							convey.So(node.Phase, convey.ShouldEqual, v1alpha1.NodePhaseNew)
						}
					}
				})

				convey.Convey("Sensor pod and service should be created", func() {
					sensorDeployment, err := controller.kubeClientset.CoreV1().Pods(soc.s.Namespace).Get("artifact-sensor", metav1.GetOptions{})
					convey.So(err, convey.ShouldBeNil)
					convey.So(sensorDeployment, convey.ShouldNotBeNil)

					sensorSvc, err := controller.kubeClientset.CoreV1().Services(sensor.Namespace).Get("artifact-sensor-svc", metav1.GetOptions{})
					convey.So(err, convey.ShouldBeNil)
					convey.So(sensorSvc, convey.ShouldNotBeNil)

					convey.Convey("Go to active state", func() {
						sensor, err := controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Get(sensor.Name, metav1.GetOptions{})
						convey.So(err, convey.ShouldBeNil)
						convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
					})
				})
			})

			convey.Convey("Operate on sensor in active state", func() {
				// Operate it once to create pod and service
				soc.markSensorPhase(v1alpha1.NodePhaseNew, false, "test")
				err := soc.operate()
				convey.So(err, convey.ShouldBeNil)
				cache.WaitForCacheSync(done, controller.podInformer.Informer().HasSynced)
				cache.WaitForCacheSync(done, controller.svcInformer.Informer().HasSynced)
				soc.markSensorPhase(v1alpha1.NodePhaseActive, false, "test")

				convey.Convey("Operation must succeed", func() {
					err := soc.operate()
					convey.So(err, convey.ShouldBeNil)

					convey.Convey("Untouch pod and service", func() {
						sensorPod, err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Get("artifact-sensor", metav1.GetOptions{})
						convey.So(err, convey.ShouldBeNil)
						convey.So(sensorPod, convey.ShouldNotBeNil)

						sensorSvc, err := controller.kubeClientset.CoreV1().Services(sensor.Namespace).Get("artifact-sensor-svc", metav1.GetOptions{})
						convey.So(err, convey.ShouldBeNil)
						convey.So(sensorSvc, convey.ShouldNotBeNil)

						convey.Convey("Stay in active state", func() {
							sensor, err := controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Get(sensor.Name, metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
						})
					})
				})

				convey.Convey("Delete pod and service", func() {
					err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Delete("artifact-sensor", &metav1.DeleteOptions{})
					convey.So(err, convey.ShouldBeNil)

					err = controller.kubeClientset.CoreV1().Services(sensor.Namespace).Delete("artifact-sensor-svc", &metav1.DeleteOptions{})
					convey.So(err, convey.ShouldBeNil)

					cache.WaitForCacheSync(done, controller.podInformer.Informer().HasSynced)
					cache.WaitForCacheSync(done, controller.svcInformer.Informer().HasSynced)

					convey.Convey("Operation must succeed", func() {
						err := soc.operate()
						convey.So(err, convey.ShouldBeNil)

						convey.Convey("Create pod and service", func() {
							sensorPod, err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Get("artifact-sensor", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorPod, convey.ShouldNotBeNil)

							sensorSvc, err := controller.kubeClientset.CoreV1().Services(sensor.Namespace).Get("artifact-sensor-svc", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorSvc, convey.ShouldNotBeNil)

							convey.Convey("Stay in active state", func() {
								sensor, err := controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Get(sensor.Name, metav1.GetOptions{})
								convey.So(err, convey.ShouldBeNil)
								convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
							})
						})
					})
				})

				convey.Convey("Change pod and service spec", func() {
					soc.srctx.s.Spec.DeploySpec.RestartPolicy = "Never"
					soc.srctx.s.Spec.EventProtocol.Http.Port = "1234"

					convey.Convey("Operation must succeed", func() {
						err := soc.operate()
						convey.So(err, convey.ShouldBeNil)

						convey.Convey("Delete pod and service", func() {
							sensorPod, err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Get("artifact-sensor", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorPod.Spec.RestartPolicy, convey.ShouldEqual, "Never")

							sensorSvc, err := controller.kubeClientset.CoreV1().Services(sensor.Namespace).Get("artifact-sensor-svc", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorSvc.Spec.Ports[0].TargetPort.IntVal, convey.ShouldEqual, 1234)

							convey.Convey("Stay in active state", func() {
								sensor, err := controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Get(sensor.Name, metav1.GetOptions{})
								convey.So(err, convey.ShouldBeNil)
								convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
							})
						})
					})
				})
			})

			convey.Convey("Operate on sensor in error state", func() {
				// Operate it once to create pod and service
				soc.markSensorPhase(v1alpha1.NodePhaseNew, false, "test")
				err := soc.operate()
				convey.So(err, convey.ShouldBeNil)
				cache.WaitForCacheSync(done, controller.podInformer.Informer().HasSynced)
				cache.WaitForCacheSync(done, controller.svcInformer.Informer().HasSynced)
				soc.markSensorPhase(v1alpha1.NodePhaseError, false, "test")

				convey.Convey("Operation must succeed", func() {
					err := soc.operate()
					convey.So(err, convey.ShouldBeNil)

					convey.Convey("Untouch pod and service", func() {
						sensorPod, err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Get("artifact-sensor", metav1.GetOptions{})
						convey.So(err, convey.ShouldBeNil)
						convey.So(sensorPod, convey.ShouldNotBeNil)

						sensorSvc, err := controller.kubeClientset.CoreV1().Services(sensor.Namespace).Get("artifact-sensor-svc", metav1.GetOptions{})
						convey.So(err, convey.ShouldBeNil)
						convey.So(sensorSvc, convey.ShouldNotBeNil)

						convey.Convey("Stay in error state", func() {
							sensor, err := controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Get(sensor.Name, metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseError)
						})
					})
				})

				convey.Convey("Delete pod and service", func() {
					err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Delete("artifact-sensor", &metav1.DeleteOptions{})
					convey.So(err, convey.ShouldBeNil)

					err = controller.kubeClientset.CoreV1().Services(sensor.Namespace).Delete("artifact-sensor-svc", &metav1.DeleteOptions{})
					convey.So(err, convey.ShouldBeNil)

					cache.WaitForCacheSync(done, controller.podInformer.Informer().HasSynced)
					cache.WaitForCacheSync(done, controller.svcInformer.Informer().HasSynced)

					convey.Convey("Operation must succeed", func() {
						err := soc.operate()
						convey.So(err, convey.ShouldBeNil)

						convey.Convey("Create pod and service", func() {
							sensorPod, err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Get("artifact-sensor", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorPod, convey.ShouldNotBeNil)

							sensorSvc, err := controller.kubeClientset.CoreV1().Services(sensor.Namespace).Get("artifact-sensor-svc", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorSvc, convey.ShouldNotBeNil)

							convey.Convey("Go to active state", func() {
								sensor, err := controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Get(sensor.Name, metav1.GetOptions{})
								convey.So(err, convey.ShouldBeNil)
								convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
							})
						})
					})
				})

				convey.Convey("Change pod and service spec", func() {
					soc.srctx.s.Spec.DeploySpec.RestartPolicy = "Never"
					soc.srctx.s.Spec.EventProtocol.Http.Port = "1234"

					convey.Convey("Operation must succeed", func() {
						err := soc.operate()
						convey.So(err, convey.ShouldBeNil)

						convey.Convey("Delete pod and service", func() {
							sensorPod, err := controller.kubeClientset.CoreV1().Pods(sensor.Namespace).Get("artifact-sensor", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorPod.Spec.RestartPolicy, convey.ShouldEqual, "Never")

							sensorSvc, err := controller.kubeClientset.CoreV1().Services(sensor.Namespace).Get("artifact-sensor-svc", metav1.GetOptions{})
							convey.So(err, convey.ShouldBeNil)
							convey.So(sensorSvc.Spec.Ports[0].TargetPort.IntVal, convey.ShouldEqual, 1234)

							convey.Convey("Go to active state", func() {
								sensor, err := controller.sensorClientset.ArgoprojV1alpha1().Sensors(sensor.Namespace).Get(sensor.Name, metav1.GetOptions{})
								convey.So(err, convey.ShouldBeNil)
								convey.So(sensor.Status.Phase, convey.ShouldEqual, v1alpha1.NodePhaseActive)
							})
						})
					})
				})
			})
		})
	})
}
