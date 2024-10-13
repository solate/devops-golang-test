/*
Copyright 2024.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MyStatefulSetSpec defines the desired state of MyStatefulSet
type MyStatefulSetSpec struct {
	// // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// // Important: Run "make" to regenerate code after modifying this file

	// // Foo is an example field of MyStatefulSet. Edit mystatefulset_types.go to remove/update
	// Foo string `json:"foo,omitempty"`

	Replicas int32                  `json:"replicas"`
	Template corev1.PodTemplateSpec `json:"template"`
}

// MyStatefulSetStatus defines the observed state of MyStatefulSet
type MyStatefulSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ReadyReplicas int32 `json:"readyReplicas"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.replicas",description="The number of replicas in the MyStatefulSet"
// +kubebuilder:printcolumn:name="Ready",type="integer",JSONPath=".status.readyReplicas",description="The number of ready replicas in the MyStatefulSet"

// MyStatefulSet is the Schema for the mystatefulsets API
type MyStatefulSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyStatefulSetSpec   `json:"spec,omitempty"`
	Status MyStatefulSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyStatefulSetList contains a list of MyStatefulSet
type MyStatefulSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyStatefulSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MyStatefulSet{}, &MyStatefulSetList{})
}
