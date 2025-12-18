// Copyright (c) ZStack.io, Inc.

package param

// DeletePciDeviceDetailParam DeletePciDevice detail param
type DeletePciDeviceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePciDeviceParam DeletePciDevice request param
type DeletePciDeviceParam struct {
	BaseParam
	Params DeletePciDeviceDetailParam `json:"params"`
}
