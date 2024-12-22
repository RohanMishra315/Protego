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

// ProtegoPolicySpec defines the desired state of ProtegoPolicy.
type ProtegoPolicySpec struct {
	// ProtegoRules is a list of rules that define the policy.
	ProtegoRules []Rule `json:"rules"`

	// Selector specifies the workload resources that the policy applies to.
	Selector WorkloadSelector `json:"selector"`
}

// ProtegoPolicyStatus defines the observed state of ProtegoPolicy.
type ProtegoPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ProtegoPolicy is the Schema for the protegopolicies API.
type ProtegoPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProtegoPolicySpec   `json:"spec,omitempty"`
	Status ProtegoPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtegoPolicyList contains a list of ProtegoPolicy.
type ProtegoPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtegoPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProtegoPolicy{}, &ProtegoPolicyList{})
}

// Rule defines a single rule within a ProtegoPolicySpec
type Rule struct {
	// ID is a unique identifier for the rule, used by security engine adapters.
	ID string `json:"id"`

	// RuleAction specifies the action to be taken when the rule matches.
	RuleAction string `json:"action"`

	// Params is an optional map of parameters associated with the rule.
	Params map[string][]string `json:"params,omitempty"`
}
