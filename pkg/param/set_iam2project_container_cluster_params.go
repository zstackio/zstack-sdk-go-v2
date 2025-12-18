// Copyright (c) ZStack.io, Inc.

package param

// SetIAM2ProjectContainerClusterDetailParam SetIAM2ProjectContainerCluster详细参数
type SetIAM2ProjectContainerClusterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"containerUuid" validate:"required"` // 必填
	rest int64 `json:"clusterId" validate:"required"` // 必填
}

// SetIAM2ProjectContainerClusterParam SetIAM2ProjectContainerCluster请求参数
type SetIAM2ProjectContainerClusterParam struct {
	BaseParam
	Params SetIAM2ProjectContainerClusterDetailParam `json:"params"` // 详细参数
}

