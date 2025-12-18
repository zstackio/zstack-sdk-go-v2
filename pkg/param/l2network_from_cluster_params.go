// Copyright (c) ZStack.io, Inc.

package param

// DetachL2NetworkFromClusterDetailParam DetachL2NetworkFromCluster详细参数
type DetachL2NetworkFromClusterDetailParam struct {
	rest string `json:"l2NetworkUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// DetachL2NetworkFromClusterParam DetachL2NetworkFromCluster请求参数
type DetachL2NetworkFromClusterParam struct {
	BaseParam
	Params DetachL2NetworkFromClusterDetailParam `json:"params"` // 详细参数
}

