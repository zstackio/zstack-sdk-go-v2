// Copyright (c) ZStack.io, Inc.

package param

// DetachPrimaryStorageFromClusterDetailParam DetachPrimaryStorageFromCluster详细参数
type DetachPrimaryStorageFromClusterDetailParam struct {
	rest string `json:"primaryStorageUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// DetachPrimaryStorageFromClusterParam DetachPrimaryStorageFromCluster请求参数
type DetachPrimaryStorageFromClusterParam struct {
	BaseParam
	Params DetachPrimaryStorageFromClusterDetailParam `json:"params"` // 详细参数
}

