// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateBareMetal2ProvisionNetworkParamDetail UpdateBareMetal2ProvisionNetwork detail param
type UpdateBareMetal2ProvisionNetworkParamDetail struct {
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
	Params UpdateBareMetal2ProvisionNetworkParamDetail `json:"params"`
}
// CreateBareMetal2ProvisionNetworkParamDetail CreateBareMetal2ProvisionNetwork detail param
type CreateBareMetal2ProvisionNetworkParamDetail struct {
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
	Params CreateBareMetal2ProvisionNetworkParamDetail `json:"params"`
}
// DeleteBareMetal2ProvisionNetworkParamDetail DeleteBareMetal2ProvisionNetwork detail param
type DeleteBareMetal2ProvisionNetworkParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBareMetal2ProvisionNetworkParam DeleteBareMetal2ProvisionNetwork request param
type DeleteBareMetal2ProvisionNetworkParam struct {
	BaseParam
	Params DeleteBareMetal2ProvisionNetworkParamDetail `json:"params"`
}
