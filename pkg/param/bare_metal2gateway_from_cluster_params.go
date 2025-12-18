// Copyright (c) ZStack.io, Inc.

package param

// DetachBareMetal2GatewayFromClusterDetailParam DetachBareMetal2GatewayFromCluster详细参数
type DetachBareMetal2GatewayFromClusterDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"gatewayUuid" validate:"required"` // 必填
}

// DetachBareMetal2GatewayFromClusterParam DetachBareMetal2GatewayFromCluster请求参数
type DetachBareMetal2GatewayFromClusterParam struct {
	BaseParam
	Params DetachBareMetal2GatewayFromClusterDetailParam `json:"params"` // 详细参数
}

