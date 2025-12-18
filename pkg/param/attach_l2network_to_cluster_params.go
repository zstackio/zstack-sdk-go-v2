// Copyright (c) ZStack.io, Inc.

package param

// AttachL2NetworkToClusterDetailParam AttachL2NetworkToCluster detail param
type AttachL2NetworkToClusterDetailParam struct {
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	L2ProviderType string `json:"l2ProviderType,omitempty"`
}

// AttachL2NetworkToClusterParam AttachL2NetworkToCluster request param
type AttachL2NetworkToClusterParam struct {
	BaseParam
	Params AttachL2NetworkToClusterDetailParam `json:"params"`
}
