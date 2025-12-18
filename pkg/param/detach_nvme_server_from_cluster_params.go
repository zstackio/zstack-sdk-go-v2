// Copyright (c) ZStack.io, Inc.

package param

// DetachNvmeServerFromClusterDetailParam DetachNvmeServerFromCluster detail param
type DetachNvmeServerFromClusterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// DetachNvmeServerFromClusterParam DetachNvmeServerFromCluster request param
type DetachNvmeServerFromClusterParam struct {
	BaseParam
	Params DetachNvmeServerFromClusterDetailParam `json:"params"`
}
