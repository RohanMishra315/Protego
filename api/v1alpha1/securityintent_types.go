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

// SecurityIntentSpec defines the desired state of SecurityIntent.
type SecurityIntentSpec struct {
	Intent Intent `json:"intent"`
}

// SecurityIntentStatus defines the observed state of SecurityIntent.
type SecurityIntentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// SecurityIntent is the Schema for the securityintents API.
type SecurityIntent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityIntentSpec   `json:"spec,omitempty"`
	Status SecurityIntentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityIntentList contains a list of SecurityIntent.
type SecurityIntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityIntent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityIntent{}, &SecurityIntentList{})
}

// Intent defines the high-level desired intent.
type Intent struct {
	// ID is predefined in adapter ID pool.
	// Used by security engines to generate corresponding security policies.
	//+kubebuilder:validation:Pattern:="^[a-zA-Z0-9]*$"
	ID string `json:"id"`

	// Action defines how the intent will be enforced.
	// Valid actions are "Audit" and "Enforce".
	Action string `json:"action"`

	// Tags are additional metadata for categorization and grouping of intents.
	// Facilitates searching, filtering, and management of security policies.
	Tags []string `json:"tags,omitempty"`

	// Params are key-value pairs that allows fine-tuning of intents to specific
	// requirements.
	Params map[string][]string `json:"params,omitempty"`
}
