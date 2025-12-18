// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ProvisionNetworkInventoryView BareMetal2ProvisionNetwork
type BareMetal2ProvisionNetworkInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DhcpInterface string `json:"dhcpInterface,omitempty"`
	DhcpRangeStartIp string `json:"dhcpRangeStartIp,omitempty"`
	DhcpRangeEndIp string `json:"dhcpRangeEndIp,omitempty"`
	DhcpRangeNetmask string `json:"dhcpRangeNetmask,omitempty"`
	DhcpRangeGateway string `json:"dhcpRangeGateway,omitempty"`
	DhcpRangeNetworkCidr string `json:"dhcpRangeNetworkCidr,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

