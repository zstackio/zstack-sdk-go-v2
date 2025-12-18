// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunProxyVSwitchDetailParam UpdateAliyunProxyVSwitch详细参数
type UpdateAliyunProxyVSwitchDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"status,omitempty"`
	rest bool `json:"isDefault,omitempty"`
}

// UpdateAliyunProxyVSwitchParam UpdateAliyunProxyVSwitch请求参数
type UpdateAliyunProxyVSwitchParam struct {
	BaseParam
	Params UpdateAliyunProxyVSwitchDetailParam `json:"params"` // 详细参数
}

