// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunProxyVSwitchDetailParam DeleteAliyunProxyVSwitch detail param
type DeleteAliyunProxyVSwitchDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunProxyVSwitchParam DeleteAliyunProxyVSwitch request param
type DeleteAliyunProxyVSwitchParam struct {
	BaseParam
	Params DeleteAliyunProxyVSwitchDetailParam `json:"params"`
}
