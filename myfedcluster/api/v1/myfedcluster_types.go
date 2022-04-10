/*
Copyright 2022.

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
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterConditionType string

// These are valid conditions of a cluster.
const (
	// ClusterReady means the cluster is ready to accept workloads.
	ClusterReady ClusterConditionType = "Ready"
	// ClusterOffline means the cluster is temporarily down or not reachable
	ClusterOffline ClusterConditionType = "Offline"
)

type TLSValidation string

const (
	TLSAll            TLSValidation = "*"
	TLSSubjectName    TLSValidation = "SubjectName"
	TLSValidityPeriod TLSValidation = "ValidityPeriod"
)

// LocalSecretReference is a reference to a secret within the enclosing
// namespace.
type LocalSecretReference struct {
	// Name of a secret within the enclosing
	// namespace
	Name string `json:"name"`
}

type ClusterCondition struct {
	// Type of cluster condition, Ready or Offline.
	Type ClusterConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status apiv1.ConditionStatus `json:"status"`
	// Last time the condition was checked.
	LastProbeTime metav1.Time `json:"lastProbeTime"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason *string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message *string `json:"message,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MyFedClusterSpec defines the desired state of MyFedCluster
type MyFedClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The API endpoint of the member cluster. This can be a hostname,
	// hostname:port, IP or IP:port.
	APIEndpoint string `json:"apiEndpoint"`

	// CABundle contains the certificate authority information.
	// +optional
	CABundle []byte `json:"caBundle,omitempty"`

	// Name of the secret containing the token required to access the
	// member cluster. The secret needs to exist in the same namespace
	// as the control plane and should have a "token" key.
	SecretRef LocalSecretReference `json:"secretRef"`

	// DisabledTLSValidations defines a list of checks to ignore when validating
	// the TLS connection to the member cluster.  This can be any of *, SubjectName, or ValidityPeriod.
	// If * is specified, it is expected to be the only option in list.
	// +optional
	DisabledTLSValidations []TLSValidation `json:"disabledTLSValidations,omitempty"`
}

// MyFedClusterStatus defines the observed state of MyFedCluster
type MyFedClusterStatus struct {
	// Conditions is an array of current cluster conditions.
	Conditions []ClusterCondition `json:"conditions"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MyFedCluster is the Schema for the myfedclusters API
type MyFedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyFedClusterSpec   `json:"spec,omitempty"`
	Status MyFedClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyFedClusterList contains a list of MyFedCluster
type MyFedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyFedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MyFedCluster{}, &MyFedClusterList{})
}
