// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2ProvisionNetworkDetailParam UpdateBareMetal2ProvisionNetwork detail param
type UpdateBareMetal2ProvisionNetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DhcpInterface string `json:"dhcpInterface,omitempty"`
	DhcpRangeStartIp string `json:"dhcpRangeStartIp,omitempty"`
	DhcpRangeEndIp string `json:"dhcpRangeEndIp,omitempty"`
	DhcpRangeNetmask string `json:"dhcpRangeNetmask,omitempty"`
	DhcpRangeGateway string `json:"dhcpRangeGateway,omitempty"`
}

// UpdateBareMetal2ProvisionNetworkParam UpdateBareMetal2ProvisionNetwork request param
type UpdateBareMetal2ProvisionNetworkParam struct {
	BaseParam
	Params UpdateBareMetal2ProvisionNetworkDetailParam `json:"params"`
}
