// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostNetworkBondingServiceRefInventoryView HostNetworkBondingServiceRef
type HostNetworkBondingServiceRefInventoryView struct {
	BondingUuid string `json:"bondingUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

