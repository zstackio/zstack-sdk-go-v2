// Copyright (c) ZStack.io, Inc.

package param

// DetachUsbDeviceFromVmDetailParam DetachUsbDeviceFromVm detail param
type DetachUsbDeviceFromVmDetailParam struct {
	UsbDeviceUuid string `json:"usbDeviceUuid" validate:"required"`
}

// DetachUsbDeviceFromVmParam DetachUsbDeviceFromVm request param
type DetachUsbDeviceFromVmParam struct {
	BaseParam
	Params DetachUsbDeviceFromVmDetailParam `json:"params"`
}
