// Copyright (c) ZStack.io, Inc.

package param

// AttachUsbDeviceToVmDetailParam AttachUsbDeviceToVm detail param
type AttachUsbDeviceToVmDetailParam struct {
	UsbDeviceUuid string `json:"usbDeviceUuid" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	AttachType string `json:"attachType,omitempty"`
}

// AttachUsbDeviceToVmParam AttachUsbDeviceToVm request param
type AttachUsbDeviceToVmParam struct {
	BaseParam
	Params AttachUsbDeviceToVmDetailParam `json:"params"`
}
