// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2GatewayStateDetailParam ChangeBareMetal2GatewayState详细参数
type ChangeBareMetal2GatewayStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeBareMetal2GatewayStateParam ChangeBareMetal2GatewayState请求参数
type ChangeBareMetal2GatewayStateParam struct {
	BaseParam
	Params ChangeBareMetal2GatewayStateDetailParam `json:"params"` // 详细参数
}

