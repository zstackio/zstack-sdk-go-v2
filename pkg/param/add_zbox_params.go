// Copyright (c) ZStack.io, Inc.

package param

// AddZBoxDetailParam AddZBox detail param
type AddZBoxDetailParam struct {
	UsbDeviceUuid string `json:"usbDeviceUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	SkipFormat bool `json:"skipFormat,omitempty"`
}

// AddZBoxParam AddZBox request param
type AddZBoxParam struct {
	BaseParam
	Params AddZBoxDetailParam `json:"params"`
}
