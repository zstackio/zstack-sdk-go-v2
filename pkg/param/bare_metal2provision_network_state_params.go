// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2ProvisionNetworkStateDetailParam ChangeBareMetal2ProvisionNetworkState详细参数
type ChangeBareMetal2ProvisionNetworkStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeBareMetal2ProvisionNetworkStateParam ChangeBareMetal2ProvisionNetworkState请求参数
type ChangeBareMetal2ProvisionNetworkStateParam struct {
	BaseParam
	Params ChangeBareMetal2ProvisionNetworkStateDetailParam `json:"params"` // 详细参数
}

