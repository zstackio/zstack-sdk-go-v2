// Copyright (c) ZStack.io, Inc.

package param

// AttachPciDeviceToVmDetailParam AttachPciDeviceToVm详细参数
type AttachPciDeviceToVmDetailParam struct {
	rest string `json:"pciDeviceUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// AttachPciDeviceToVmParam AttachPciDeviceToVm请求参数
type AttachPciDeviceToVmParam struct {
	BaseParam
	Params AttachPciDeviceToVmDetailParam `json:"params"` // 详细参数
}

