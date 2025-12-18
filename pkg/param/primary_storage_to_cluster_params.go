// Copyright (c) ZStack.io, Inc.

package param

// AttachPrimaryStorageToClusterDetailParam AttachPrimaryStorageToCluster详细参数
type AttachPrimaryStorageToClusterDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
}

// AttachPrimaryStorageToClusterParam AttachPrimaryStorageToCluster请求参数
type AttachPrimaryStorageToClusterParam struct {
	BaseParam
	Params AttachPrimaryStorageToClusterDetailParam `json:"params"` // 详细参数
}

