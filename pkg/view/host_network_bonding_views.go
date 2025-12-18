// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostNetworkBondingInventoryView HostNetworkBonding
type HostNetworkBondingInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"bondingName,omitempty"`
	rest string `json:"bondingType,omitempty"`
	rest int64 `json:"speed,omitempty"`
	rest string `json:"mode,omitempty"`
	rest string `json:"xmitHashPolicy,omitempty"`
	rest string `json:"miiStatus,omitempty"`
	rest string `json:"mac,omitempty"`
	rest []string `json:"ipAddresses,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"callBackIp,omitempty"`
	rest int64 `json:"miimon,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"allSlavesActive,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []HostNetworkInterfaceInventoryView `json:"slaves,omitempty"`
}

