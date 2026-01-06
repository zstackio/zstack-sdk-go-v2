// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IpRangeInventoryView IpRange
type IpRangeInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StartIp string `json:"startIp,omitempty"`
	EndIp string `json:"endIp,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	NetworkCidr string `json:"networkCidr,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	AddressMode string `json:"addressMode,omitempty"`
	PrefixLen int `json:"prefixLen,omitempty"`
	IpRangeType string `json:"ipRangeType,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// AddIpRangeEventView AddIpRangeEvent
type AddIpRangeEventView struct {
	Inventory IpRangeInventoryView `json:"inventory,omitempty"`
}

// AddIpRangeByNetworkCidrEventView AddIpRangeByNetworkCidrEvent
type AddIpRangeByNetworkCidrEventView struct {
	Inventory IpRangeInventoryView `json:"inventory,omitempty"`
	Inventories []IpRangeInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UpdateIpRangeEventView UpdateIpRangeEvent
type UpdateIpRangeEventView struct {
	Inventory IpRangeInventoryView `json:"inventory,omitempty"`
}

// QueryIpRangeView QueryIpRange
type QueryIpRangeView struct {
	Inventories []IpRangeInventoryView `json:"inventories,omitempty"`
}

// DeleteIpRangeEventView DeleteIpRangeEvent
type DeleteIpRangeEventView struct {
	Success bool `json:"success,omitempty"`
}

