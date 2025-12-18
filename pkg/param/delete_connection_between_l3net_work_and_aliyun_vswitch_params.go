// Copyright (c) ZStack.io, Inc.

package param

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch detail param
type DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParam DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch request param
type DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchParam struct {
	BaseParam
	Params DeleteConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam `json:"params"`
}
