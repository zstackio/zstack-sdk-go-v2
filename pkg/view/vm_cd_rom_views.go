// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmCdRomInventoryView VmCdRom
type VmCdRomInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	IsoUuid string `json:"isoUuid,omitempty"`
	IsoInstallPath string `json:"isoInstallPath,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// CreateVmCdRomEventView CreateVmCdRomEvent
type CreateVmCdRomEventView struct {
	Inventory VmCdRomInventoryView `json:"inventory,omitempty"`
}

// QueryVmCdRomView QueryVmCdRom
type QueryVmCdRomView struct {
	Inventories []VmCdRomInventoryView `json:"inventories,omitempty"`
}

// UpdateVmCdRomEventView UpdateVmCdRomEvent
type UpdateVmCdRomEventView struct {
	Inventory VmCdRomInventoryView `json:"inventory,omitempty"`
}

// SetVmInstanceDefaultCdRomEventView SetVmInstanceDefaultCdRomEvent
type SetVmInstanceDefaultCdRomEventView struct {
	Inventory VmCdRomInventoryView `json:"inventory,omitempty"`
}

