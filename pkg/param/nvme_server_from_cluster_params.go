// Copyright (c) ZStack.io, Inc.

package param

// DetachNvmeServerFromClusterDetailParam DetachNvmeServerFromCluster详细参数
type DetachNvmeServerFromClusterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// DetachNvmeServerFromClusterParam DetachNvmeServerFromCluster请求参数
type DetachNvmeServerFromClusterParam struct {
	BaseParam
	Params DetachNvmeServerFromClusterDetailParam `json:"params"` // 详细参数
}

