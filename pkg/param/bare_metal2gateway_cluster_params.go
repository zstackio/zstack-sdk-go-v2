// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2GatewayClusterDetailParam ChangeBareMetal2GatewayCluster详细参数
type ChangeBareMetal2GatewayClusterDetailParam struct {
	rest string `json:"gatewayUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// ChangeBareMetal2GatewayClusterParam ChangeBareMetal2GatewayCluster请求参数
type ChangeBareMetal2GatewayClusterParam struct {
	BaseParam
	Params ChangeBareMetal2GatewayClusterDetailParam `json:"params"` // 详细参数
}

