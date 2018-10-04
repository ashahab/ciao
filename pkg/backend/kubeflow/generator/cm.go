// Copyright 2018 Caicloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"fmt"
	"os"
	s2iconfigmap "github.com/caicloud/ciao/pkg/s2i/configmap"
	pytorchv1alpha2 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1alpha2"
	tfv1alpha2 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"github.com/caicloud/ciao/pkg/types"
)

const (
	baseImagePyTorch = "pytorch/pytorch:v0.2"
)

// CM is the type for CM generator.
type CM struct {
}

// NewCM returns a new CM generator.
func NewCM() *CM {
	return &CM{}
}


// GenerateTFJob generates a new TFJob.
func (c CM) GenerateTFJob(parameter *types.Parameter) *tfv1alpha2.TFJob {
	psCount := int32(parameter.PSCount)
	workerCount := int32(parameter.WorkerCount)
	cheifCount := int32(1)
	evalCount := int32(1)
	mountPath := fmt.Sprintf("/%s", parameter.Image)
	filename := fmt.Sprintf("/%s/%s", parameter.Image, s2iconfigmap.FileName)
	image_name := os.Getenv("IMAGE_NAME")
	cpu_image_name := os.Getenv("CPU_IMAGE_NAME")
	var gpuResourceName v1.ResourceName
	gpuResourceName = "nvidia.com/gpu"
	return &tfv1alpha2.TFJob{
		TypeMeta: metav1.TypeMeta{
			Kind: tfv1alpha2.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      parameter.GenerateName,
			Namespace: parameter.Namespace,
		},
		Spec: tfv1alpha2.TFJobSpec{
			TFReplicaSpecs: map[tfv1alpha2.TFReplicaType]*tfv1alpha2.TFReplicaSpec{
				tfv1alpha2.TFReplicaTypePS: &tfv1alpha2.TFReplicaSpec{
					Replicas: &psCount,
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							NodeSelector: map[string]string{
								"node-role.kubernetes.io/cpu": "cpu",
							},
							Containers: []v1.Container{
								v1.Container{
									Name:  defaultContainerNameTF,
									Image: cpu_image_name,
									Command: []string{
										"python",
										filename,
									},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      parameter.Image,
											MountPath: mountPath,
										},
									},
								},
							},
							Volumes: []v1.Volume{
								v1.Volume{
									Name: parameter.Image,
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: parameter.Image,
											},
										},
									},
								},
							},
						},
					},
				},
				tfv1alpha2.TFReplicaTypeWorker: &tfv1alpha2.TFReplicaSpec{
					Replicas: &workerCount,
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  defaultContainerNameTF,
									Image: image_name,
									Resources: v1.ResourceRequirements{
										Limits: v1.ResourceList{
											gpuResourceName: *resource.NewQuantity(1, resource.DecimalSI),
										},
									},
									Command: []string{
										"python",
										filename,
									},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      parameter.Image,
											MountPath: mountPath,
										},
									},
								},
							},
							Volumes: []v1.Volume{
								v1.Volume{
									Name: parameter.Image,
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: parameter.Image,
											},
										},
									},
								},
							},
						},
					},
				},
				tfv1alpha2.TFReplicaTypeChief: &tfv1alpha2.TFReplicaSpec{
					Replicas: &cheifCount,
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  defaultContainerNameTF,
									Image: image_name,
									Resources: v1.ResourceRequirements {
										Limits: v1.ResourceList{
											gpuResourceName: *resource.NewQuantity(1, resource.DecimalSI),
										},
									},
									Command: []string{
										"python",
										filename,
									},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      parameter.Image,
											MountPath: mountPath,
										},
									},
								},
							},
							Volumes: []v1.Volume{
								v1.Volume{
									Name: parameter.Image,
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: parameter.Image,
											},
										},
									},
								},
							},
						},
					},
				},
				tfv1alpha2.TFReplicaTypeEval: &tfv1alpha2.TFReplicaSpec{
					Replicas: &evalCount,
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  defaultContainerNameTF,
									Image: image_name,
									Resources: v1.ResourceRequirements{
										Limits: v1.ResourceList{
											gpuResourceName: *resource.NewQuantity(1, resource.DecimalSI),
										},
									},
									Command: []string{
										"python",
										filename,
									},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      parameter.Image,
											MountPath: mountPath,
										},
									},
								},
							},
							Volumes: []v1.Volume{
								v1.Volume{
									Name: parameter.Image,
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: parameter.Image,
											},
										},
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

// GeneratePyTorchJob generates a new PyTorchJob.
func (c CM) GeneratePyTorchJob(parameter *types.Parameter) *pytorchv1alpha2.PyTorchJob {
	masterCount := int32(parameter.MasterCount)
	workerCount := int32(parameter.WorkerCount)

	mountPath := fmt.Sprintf("/%s", parameter.Image)
	filename := fmt.Sprintf("/%s/%s", parameter.Image, s2iconfigmap.FileName)
	return &pytorchv1alpha2.PyTorchJob{
		TypeMeta: metav1.TypeMeta{
			Kind: pytorchv1alpha2.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      parameter.GenerateName,
			Namespace: metav1.NamespaceDefault,
		},
		Spec: pytorchv1alpha2.PyTorchJobSpec{
			PyTorchReplicaSpecs: map[pytorchv1alpha2.PyTorchReplicaType]*pytorchv1alpha2.PyTorchReplicaSpec{
				pytorchv1alpha2.PyTorchReplicaTypeMaster: &pytorchv1alpha2.PyTorchReplicaSpec{
					Replicas: &masterCount,
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  defaultContainerNamePyTorch,
									Image: baseImagePyTorch,
									Command: []string{
										"python",
										filename,
									},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      parameter.Image,
											MountPath: mountPath,
										},
									},
								},
							},
							Volumes: []v1.Volume{
								v1.Volume{
									Name: parameter.Image,
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: parameter.Image,
											},
										},
									},
								},
							},
						},
					},
				},
				pytorchv1alpha2.PyTorchReplicaTypeWorker: &pytorchv1alpha2.PyTorchReplicaSpec{
					Replicas: &workerCount,
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  defaultContainerNamePyTorch,
									Image: baseImagePyTorch,
									Command: []string{
										"python",
										filename,
									},
									VolumeMounts: []v1.VolumeMount{
										v1.VolumeMount{
											Name:      parameter.Image,
											MountPath: mountPath,
										},
									},
								},
							},
							Volumes: []v1.Volume{
								v1.Volume{
									Name: parameter.Image,
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: parameter.Image,
											},
										},
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
