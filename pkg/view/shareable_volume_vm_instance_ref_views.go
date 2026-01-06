// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ShareableVolumeVmInstanceRefInventoryView ShareableVolumeVmInstanceRef
type ShareableVolumeVmInstanceRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryShareableVolumeVmInstanceRefView QueryShareableVolumeVmInstanceRef
type QueryShareableVolumeVmInstanceRefView struct {
	Inventories []ShareableVolumeVmInstanceRefInventoryView `json:"inventories,omitempty"`
}

