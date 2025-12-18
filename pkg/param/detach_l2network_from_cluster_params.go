// Copyright (c) ZStack.io, Inc.

package param

// DetachL2NetworkFromClusterDetailParam DetachL2NetworkFromCluster detail param
type DetachL2NetworkFromClusterDetailParam struct {
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// DetachL2NetworkFromClusterParam DetachL2NetworkFromCluster request param
type DetachL2NetworkFromClusterParam struct {
	BaseParam
	Params DetachL2NetworkFromClusterDetailParam `json:"params"`
}
