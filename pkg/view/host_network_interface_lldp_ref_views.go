// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostNetworkInterfaceLldpRefInventoryView HostNetworkInterfaceLldpRef
type HostNetworkInterfaceLldpRefInventoryView struct {
	rest string `json:"lldpUuid,omitempty"`
	rest string `json:"chassisId,omitempty"`
	rest int `json:"timeToLive,omitempty"`
	rest string `json:"managementAddress,omitempty"`
	rest string `json:"systemName,omitempty"`
	rest string `json:"systemDescription,omitempty"`
	rest string `json:"systemCapabilities,omitempty"`
	rest string `json:"portId,omitempty"`
	rest string `json:"portDescription,omitempty"`
	rest int `json:"vlanId,omitempty"`
	rest int64 `json:"aggregationPortId,omitempty"`
	rest int `json:"mtu,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

