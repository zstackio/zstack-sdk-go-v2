// Copyright (c) ZStack.io, Inc.

package param

// DeleteBareMetal2ProvisionNetworkDetailParam DeleteBareMetal2ProvisionNetwork详细参数
type DeleteBareMetal2ProvisionNetworkDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteBareMetal2ProvisionNetworkParam DeleteBareMetal2ProvisionNetwork请求参数
type DeleteBareMetal2ProvisionNetworkParam struct {
	BaseParam
	Params DeleteBareMetal2ProvisionNetworkDetailParam `json:"params"` // 详细参数
}

