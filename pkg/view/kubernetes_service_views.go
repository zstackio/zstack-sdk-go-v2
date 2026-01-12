// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// KubernetesServiceInventoryView KubernetesService
type KubernetesServiceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description *string `json:"description,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Type *string `json:"type,omitempty"`
	ClusterIp *string `json:"clusterIp,omitempty"`
	ExternalIp *string `json:"externalIp,omitempty"`
	Ports *string `json:"ports,omitempty"`
	EndpointUuid *string `json:"endpointUuid,omitempty"`
	ClusterId *int64 `json:"clusterId,omitempty"`
}

