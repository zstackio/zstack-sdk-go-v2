// Copyright (c) ZStack.io, Inc.

package param

// DeletePciDeviceOfferingDetailParam DeletePciDeviceOffering detail param
type DeletePciDeviceOfferingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePciDeviceOfferingParam DeletePciDeviceOffering request param
type DeletePciDeviceOfferingParam struct {
	BaseParam
	Params DeletePciDeviceOfferingDetailParam `json:"params"`
}
