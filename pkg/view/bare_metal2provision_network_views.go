// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2ProvisionNetworkInventoryView BareMetal2ProvisionNetwork
type BareMetal2ProvisionNetworkInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"dhcpInterface,omitempty"`
	rest string `json:"dhcpRangeStartIp,omitempty"`
	rest string `json:"dhcpRangeEndIp,omitempty"`
	rest string `json:"dhcpRangeNetmask,omitempty"`
	rest string `json:"dhcpRangeGateway,omitempty"`
	rest string `json:"dhcpRangeNetworkCidr,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"attachedClusterUuids,omitempty"`
}

