// Copyright (c) ZStack.io, Inc.

package param

// DetachPciDeviceFromVmDetailParam DetachPciDeviceFromVm详细参数
type DetachPciDeviceFromVmDetailParam struct {
	rest string `json:"pciDeviceUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// DetachPciDeviceFromVmParam DetachPciDeviceFromVm请求参数
type DetachPciDeviceFromVmParam struct {
	BaseParam
	Params DetachPciDeviceFromVmDetailParam `json:"params"` // 详细参数
}

