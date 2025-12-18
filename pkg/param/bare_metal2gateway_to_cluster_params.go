// Copyright (c) ZStack.io, Inc.

package param

// AttachBareMetal2GatewayToClusterDetailParam AttachBareMetal2GatewayToCluster详细参数
type AttachBareMetal2GatewayToClusterDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"gatewayUuid" validate:"required"` // 必填
}

// AttachBareMetal2GatewayToClusterParam AttachBareMetal2GatewayToCluster请求参数
type AttachBareMetal2GatewayToClusterParam struct {
	BaseParam
	Params AttachBareMetal2GatewayToClusterDetailParam `json:"params"` // 详细参数
}

