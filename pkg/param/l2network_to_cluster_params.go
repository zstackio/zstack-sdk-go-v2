// Copyright (c) ZStack.io, Inc.

package param

// AttachL2NetworkToClusterDetailParam AttachL2NetworkToCluster详细参数
type AttachL2NetworkToClusterDetailParam struct {
	rest string `json:"l2NetworkUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"l2ProviderType,omitempty"`
}

// AttachL2NetworkToClusterParam AttachL2NetworkToCluster请求参数
type AttachL2NetworkToClusterParam struct {
	BaseParam
	Params AttachL2NetworkToClusterDetailParam `json:"params"` // 详细参数
}

