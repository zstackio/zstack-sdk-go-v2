// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ScsiLunVmInstanceRefInventoryView ScsiLunVmInstanceRef
type ScsiLunVmInstanceRefInventoryView struct {
	ScsiLunUuid string `json:"scsiLunUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	AttachMultipath bool `json:"attachMultipath,omitempty"`
}

