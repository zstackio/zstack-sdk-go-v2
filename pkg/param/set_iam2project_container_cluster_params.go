// Copyright (c) ZStack.io, Inc.

package param

// SetIAM2ProjectContainerClusterDetailParam SetIAM2ProjectContainerCluster detail param
type SetIAM2ProjectContainerClusterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ContainerUuid string `json:"containerUuid" validate:"required"`
	ClusterId int64 `json:"clusterId" validate:"required"`
}

// SetIAM2ProjectContainerClusterParam SetIAM2ProjectContainerCluster request param
type SetIAM2ProjectContainerClusterParam struct {
	BaseParam
	Params SetIAM2ProjectContainerClusterDetailParam `json:"params"`
}
