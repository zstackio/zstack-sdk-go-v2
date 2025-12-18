// Copyright (c) ZStack.io, Inc.

package view

// AddIpRangeByNetworkCidrEventView AddIpRangeByNetworkCidrEvent
type AddIpRangeByNetworkCidrEventView struct {
	Inventory IpRangeInventoryView `json:"inventory,omitempty"`
	Inventories []IpRangeInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

