// Copyright (c) ZStack.io, Inc.

package param

// UpdatePciDeviceDetailParam UpdatePciDevice detail param
type UpdatePciDeviceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	State string `json:"state,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	MetaData string `json:"metaData,omitempty"`
}

// UpdatePciDeviceParam UpdatePciDevice request param
type UpdatePciDeviceParam struct {
	BaseParam
	Params UpdatePciDeviceDetailParam `json:"params"`
}
