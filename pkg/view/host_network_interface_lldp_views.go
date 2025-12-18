// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostNetworkInterfaceLldpInventoryView HostNetworkInterfaceLldp
type HostNetworkInterfaceLldpInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"interfaceUuid,omitempty"`
	rest string `json:"mode,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest HostNetworkInterfaceLldpRefInventoryView `json:"neighborDevice,omitempty"`
}

