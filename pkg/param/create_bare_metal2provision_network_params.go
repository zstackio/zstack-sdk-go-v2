// Copyright (c) ZStack.io, Inc.

package param

// CreateBareMetal2ProvisionNetworkDetailParam CreateBareMetal2ProvisionNetwork detail param
type CreateBareMetal2ProvisionNetworkDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	DhcpInterface string `json:"dhcpInterface" validate:"required"`
	DhcpRangeStartIp string `json:"dhcpRangeStartIp" validate:"required"`
	DhcpRangeEndIp string `json:"dhcpRangeEndIp" validate:"required"`
	DhcpRangeNetmask string `json:"dhcpRangeNetmask" validate:"required"`
	DhcpRangeGateway string `json:"dhcpRangeGateway,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBareMetal2ProvisionNetworkParam CreateBareMetal2ProvisionNetwork request param
type CreateBareMetal2ProvisionNetworkParam struct {
	BaseParam
	Params CreateBareMetal2ProvisionNetworkDetailParam `json:"params"`
}
