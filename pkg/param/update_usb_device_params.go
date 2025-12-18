// Copyright (c) ZStack.io, Inc.

package param

// UpdateUsbDeviceDetailParam UpdateUsbDevice detail param
type UpdateUsbDeviceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateUsbDeviceParam UpdateUsbDevice request param
type UpdateUsbDeviceParam struct {
	BaseParam
	Params UpdateUsbDeviceDetailParam `json:"params"`
}
