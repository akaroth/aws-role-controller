/*
Copyright 2025.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AWSRoleAssignmentSpec defines the desired state of AWSRoleAssignment.
type AWSRoleAssignmentSpec struct {
	RoleName   string   `json:"roleName"`
	PolicyARNs []string `json:"policyARNs"`
}

// AWSRoleAssignmentStatus defines the observed state of AWSRoleAssignment.
type AWSRoleAssignmentStatus struct {
	Synced    bool   `json:"synced,omitempty"`
	LastError string `json:"lastError,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AWSRoleAssignment is the Schema for the awsroleassignments API.
type AWSRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSRoleAssignmentSpec   `json:"spec,omitempty"`
	Status AWSRoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSRoleAssignmentList contains a list of AWSRoleAssignment.
type AWSRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSRoleAssignment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSRoleAssignment{}, &AWSRoleAssignmentList{})
}
