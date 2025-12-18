// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostNetworkBondingInventoryView HostNetworkBonding
type HostNetworkBondingInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	BondingName string `json:"bondingName,omitempty"`
	BondingType string `json:"bondingType,omitempty"`
	Speed int64 `json:"speed,omitempty"`
	Mode string `json:"mode,omitempty"`
	XmitHashPolicy string `json:"xmitHashPolicy,omitempty"`
	MiiStatus string `json:"miiStatus,omitempty"`
	Mac string `json:"mac,omitempty"`
	IpAddresses []string `json:"ipAddresses,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	CallBackIp string `json:"callBackIp,omitempty"`
	Miimon int64 `json:"miimon,omitempty"`
	Type string `json:"type,omitempty"`
	AllSlavesActive bool `json:"allSlavesActive,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Slaves []HostNetworkInterfaceInventoryView `json:"slaves,omitempty"`
}

