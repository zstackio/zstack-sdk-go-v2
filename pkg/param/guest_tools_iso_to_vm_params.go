// Copyright (c) ZStack.io, Inc.

package param

// AttachGuestToolsIsoToVmDetailParam AttachGuestToolsIsoToVm详细参数
type AttachGuestToolsIsoToVmDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// AttachGuestToolsIsoToVmParam AttachGuestToolsIsoToVm请求参数
type AttachGuestToolsIsoToVmParam struct {
	BaseParam
	Params AttachGuestToolsIsoToVmDetailParam `json:"params"` // 详细参数
}

