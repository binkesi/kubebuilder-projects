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

package controllers

import (
	kubeclientset "k8s.io/client-go/kubernetes"
)

const (
	UserAgentName = "Cluster-Controller"

	// Following labels come from k8s.io/kubernetes/pkg/kubelet/apis
	LabelZoneFailureDomain = "failure-domain.beta.kubernetes.io/zone"
	LabelZoneRegion        = "failure-domain.beta.kubernetes.io/region"

	// Common ClusterConditions for KubeFedClusterStatus
	ClusterReady              = "ClusterReady"
	HealthzOk                 = "/healthz responded with ok"
	ClusterNotReady           = "ClusterNotReady"
	HealthzNotOk              = "/healthz responded without ok"
	ClusterNotReachableReason = "ClusterNotReachable"
	ClusterNotReachableMsg    = "cluster is not reachable"
	ClusterReachableReason    = "ClusterReachable"
	ClusterReachableMsg       = "cluster is reachable"
)

// particular KubeFedCluster.
type ClusterClient struct {
	kubeClient  *kubeclientset.Clientset
	clusterName string
}