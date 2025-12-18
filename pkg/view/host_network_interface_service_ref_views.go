// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostNetworkInterfaceServiceRefInventoryView HostNetworkInterfaceServiceRef
type HostNetworkInterfaceServiceRefInventoryView struct {
	rest string `json:"interfaceUuid,omitempty"`
	rest int `json:"vlanId,omitempty"`
	rest string `json:"serviceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

