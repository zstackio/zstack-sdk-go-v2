// Copyright (c) ZStack.io, Inc.

package param

// DetachPrimaryStorageFromClusterDetailParam DetachPrimaryStorageFromCluster detail param
type DetachPrimaryStorageFromClusterDetailParam struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// DetachPrimaryStorageFromClusterParam DetachPrimaryStorageFromCluster request param
type DetachPrimaryStorageFromClusterParam struct {
	BaseParam
	Params DetachPrimaryStorageFromClusterDetailParam `json:"params"`
}
