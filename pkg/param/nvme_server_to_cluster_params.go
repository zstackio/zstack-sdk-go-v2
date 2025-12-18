// Copyright (c) ZStack.io, Inc.

package param

// AttachNvmeServerToClusterDetailParam AttachNvmeServerToCluster详细参数
type AttachNvmeServerToClusterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// AttachNvmeServerToClusterParam AttachNvmeServerToCluster请求参数
type AttachNvmeServerToClusterParam struct {
	BaseParam
	Params AttachNvmeServerToClusterDetailParam `json:"params"` // 详细参数
}

