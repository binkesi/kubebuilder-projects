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
	"github.com/binkesi/kubebuilder-projects/myfedcluster/api/v1/common"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FedClusterSpec defines the desired state of FedCluster
type FedClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of FedCluster. Edit fedcluster_types.go to remove/update
	APIEndpoint string               `json:"apiEndpoint"`
	CABundle    []byte               `json:"caBundle,omitempty"`
	SecretRef   LocalSecretReference `json:"secretRef"`
}

// The local secret with same namespace
type LocalSecretReference struct {
	Name string `json:"name"`
}

// FedClusterStatus defines the observed state of FedCluster
type FedClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []ClusterCondition `json:"conditions"`
}

type ClusterCondition struct {
	// Type of cluster condition, Ready or Offline.
	Type common.ClusterConditionType `json:"type"`
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

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FedCluster is the Schema for the fedclusters API
type FedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FedClusterSpec   `json:"spec,omitempty"`
	Status FedClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FedClusterList contains a list of FedCluster
type FedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FedCluster{}, &FedClusterList{})
}
