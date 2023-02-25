/*
Copyright 2023 james.y.yang.

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

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DemoAppSpec defines the desired state of DemoApp
type DemoAppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Deployment DeploymentTemplate `json:"deployment,omitempty"`
	Service    ServiceTemplate    `json:"service,omitempty"`
}

type DeploymentTemplate struct {
	appsv1.DeploymentSpec `json:",inline"`
}

type ServiceTemplate struct {
	corev1.ServiceSpec `json:",inline"`
}

// DemoAppStatus defines the observed state of DemoApp
type DemoAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Workflow appsv1.DeploymentStatus `json:"workflow"`
	Network  corev1.ServiceStatus    `json:"network"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DemoApp is the Schema for the demoapps API
type DemoApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoAppSpec   `json:"spec,omitempty"`
	Status DemoAppStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DemoAppList contains a list of DemoApp
type DemoAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoApp{}, &DemoAppList{})
}
