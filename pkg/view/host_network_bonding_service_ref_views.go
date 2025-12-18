// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostNetworkBondingServiceRefInventoryView HostNetworkBondingServiceRef
type HostNetworkBondingServiceRefInventoryView struct {
	rest string `json:"bondingUuid,omitempty"`
	rest int `json:"vlanId,omitempty"`
	rest string `json:"serviceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

