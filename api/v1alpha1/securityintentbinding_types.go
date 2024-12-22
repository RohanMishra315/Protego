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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SecurityIntentBindingSpec defines the desired state of SecurityIntentBinding.
type SecurityIntentBindingSpec struct {
	Intents  []MatchIntent    `json:"intents"`
	Selector WorkloadSelector `json:"selector"`
}

// SecurityIntentBindingStatus defines the observed state of SecurityIntentBinding.
type SecurityIntentBindingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SecurityIntentBinding is the Schema for the securityintentbindings API.
type SecurityIntentBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityIntentBindingSpec   `json:"spec,omitempty"`
	Status SecurityIntentBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityIntentBindingList contains a list of SecurityIntentBinding.
type SecurityIntentBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityIntentBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityIntentBinding{}, &SecurityIntentBindingList{})
}

// MatchIntent represents an intent definition.
type MatchIntent struct {
	Name string `json:"name"`
}

// WorkloadSelector defines a selector for workloads based on labels.
type WorkloadSelector struct {
	MatchLabels map[string]string `json:"matchLabels"`
}
