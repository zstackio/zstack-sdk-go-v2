// Copyright (c) ZStack.io, Inc.

package param

// ReconnectBareMetal2GatewayDetailParam ReconnectBareMetal2Gateway详细参数
type ReconnectBareMetal2GatewayDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ReconnectBareMetal2GatewayParam ReconnectBareMetal2Gateway请求参数
type ReconnectBareMetal2GatewayParam struct {
	BaseParam
	Params ReconnectBareMetal2GatewayDetailParam `json:"params"` // 详细参数
}

