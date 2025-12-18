// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunProxyVSwitchDetailParam UpdateAliyunProxyVSwitch detail param
type UpdateAliyunProxyVSwitchDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Status string `json:"status,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// UpdateAliyunProxyVSwitchParam UpdateAliyunProxyVSwitch request param
type UpdateAliyunProxyVSwitchParam struct {
	BaseParam
	Params UpdateAliyunProxyVSwitchDetailParam `json:"params"`
}
