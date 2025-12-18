// Copyright (c) ZStack.io, Inc.

package param

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch detail param
type UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch request param
type UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam struct {
	BaseParam
	Params UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchDetailParam `json:"params"`
}
