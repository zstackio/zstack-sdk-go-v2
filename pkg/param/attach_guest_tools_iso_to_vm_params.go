// Copyright (c) ZStack.io, Inc.

package param

// AttachGuestToolsIsoToVmDetailParam AttachGuestToolsIsoToVm detail param
type AttachGuestToolsIsoToVmDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// AttachGuestToolsIsoToVmParam AttachGuestToolsIsoToVm request param
type AttachGuestToolsIsoToVmParam struct {
	BaseParam
	Params AttachGuestToolsIsoToVmDetailParam `json:"params"`
}
