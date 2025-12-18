// Copyright (c) ZStack.io, Inc.

package param

// DetachUsbDeviceFromVmDetailParam DetachUsbDeviceFromVm详细参数
type DetachUsbDeviceFromVmDetailParam struct {
	rest string `json:"usbDeviceUuid" validate:"required"` // 必填
}

// DetachUsbDeviceFromVmParam DetachUsbDeviceFromVm请求参数
type DetachUsbDeviceFromVmParam struct {
	BaseParam
	Params DetachUsbDeviceFromVmDetailParam `json:"params"` // 详细参数
}

