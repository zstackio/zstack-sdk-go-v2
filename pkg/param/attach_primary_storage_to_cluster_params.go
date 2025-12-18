// Copyright (c) ZStack.io, Inc.

package param

// AttachPrimaryStorageToClusterDetailParam AttachPrimaryStorageToCluster detail param
type AttachPrimaryStorageToClusterDetailParam struct {
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
}

// AttachPrimaryStorageToClusterParam AttachPrimaryStorageToCluster request param
type AttachPrimaryStorageToClusterParam struct {
	BaseParam
	Params AttachPrimaryStorageToClusterDetailParam `json:"params"`
}
