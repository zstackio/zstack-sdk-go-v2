// Copyright (c) ZStack.io, Inc.

package param

// UpdateBaremetalPxeServerDetailParam UpdateBaremetalPxeServer detail param
type UpdateBaremetalPxeServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DhcpRangeBegin string `json:"dhcpRangeBegin,omitempty"`
	DhcpRangeEnd string `json:"dhcpRangeEnd,omitempty"`
	DhcpRangeNetmask string `json:"dhcpRangeNetmask,omitempty"`
}

// UpdateBaremetalPxeServerParam UpdateBaremetalPxeServer request param
type UpdateBaremetalPxeServerParam struct {
	BaseParam
	Params UpdateBaremetalPxeServerDetailParam `json:"params"`
}
