// Copyright (c) ZStack.io, Inc.

package param

// AttachUsbDeviceToVmDetailParam AttachUsbDeviceToVm详细参数
type AttachUsbDeviceToVmDetailParam struct {
	rest string `json:"usbDeviceUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"attachType,omitempty"`
}

// AttachUsbDeviceToVmParam AttachUsbDeviceToVm请求参数
type AttachUsbDeviceToVmParam struct {
	BaseParam
	Params AttachUsbDeviceToVmDetailParam `json:"params"` // 详细参数
}

