// Copyright (c) ZStack.io, Inc.

package param

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch详细参数
type DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParam DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch请求参数
type DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParam struct {
	BaseParam
	Params DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam `json:"params"` // 详细参数
}

