// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmVdpaNicInventoryView VmVdpaNic
type VmVdpaNicInventoryView struct {
	PciDeviceUuid string `json:"pciDeviceUuid,omitempty"`
	LastPciDeviceUuid string `json:"lastPciDeviceUuid,omitempty"`
	SrcPath string `json:"srcPath,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Ip string `json:"ip,omitempty"`
	Mac string `json:"mac,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	MetaData string `json:"metaData,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	DriverType string `json:"driverType,omitempty"`
	UsedIps []UsedIpInventoryView `json:"usedIps,omitempty"`
	InternalName string `json:"internalName,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

