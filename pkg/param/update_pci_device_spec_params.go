// Copyright (c) ZStack.io, Inc.

package param

// UpdatePciDeviceSpecDetailParam UpdatePciDeviceSpec detail param
type UpdatePciDeviceSpecDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	RomContent string `json:"romContent,omitempty"`
	RomVersion string `json:"romVersion,omitempty"`
	AbandonSpecRom bool `json:"abandonSpecRom,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdatePciDeviceSpecParam UpdatePciDeviceSpec request param
type UpdatePciDeviceSpecParam struct {
	BaseParam
	Params UpdatePciDeviceSpecDetailParam `json:"params"`
}
