// Copyright (c) ZStack.io, Inc.

package param

// GetConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam GetConnectionBetweenL3NetworkAndAliyunVSwitch detail param
type GetConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceType string `json:"resourceType" validate:"required"`
}

// GetConnectionBetweenL3NetworkAndAliyunVSwitchParam GetConnectionBetweenL3NetworkAndAliyunVSwitch request param
type GetConnectionBetweenL3NetworkAndAliyunVSwitchParam struct {
	BaseParam
	Params GetConnectionBetweenL3NetworkAndAliyunVSwitchDetailParam `json:"params"`
}
