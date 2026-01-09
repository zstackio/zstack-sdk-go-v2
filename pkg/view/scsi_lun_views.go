// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ScsiLunInventoryView ScsiLun
type ScsiLunInventoryView struct {
	ScsiLunHostRefs []ScsiLunHostRefInventoryView `json:"scsiLunHostRefs,omitempty"`
	ScsiLunVmInstanceRefs []ScsiLunVmInstanceRefInventoryView `json:"scsiLunVmInstanceRefs,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Wwid *string `json:"wwid,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Model *string `json:"model,omitempty"`
	Wwn *string `json:"wwn,omitempty"`
	Serial *string `json:"serial,omitempty"`
	Type *string `json:"type,omitempty"`
	Hctl *string `json:"hctl,omitempty"`
	Path *string `json:"path,omitempty"`
	State *string `json:"state,omitempty"`
	Size *int64 `json:"size,omitempty"`
	MultipathDeviceUuid *string `json:"multipathDeviceUuid,omitempty"`
	Source *string `json:"source,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// UpdateScsiLunEventView UpdateScsiLunEvent
type UpdateScsiLunEventView struct {
	Inventory ScsiLunInventoryView `json:"inventory,omitempty"`
}

// GetScsiLunCandidatesForAttachingVmView GetScsiLunCandidatesForAttachingVm
type GetScsiLunCandidatesForAttachingVmView struct {
	Inventories []ScsiLunInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// QueryScsiLunView QueryScsiLun
type QueryScsiLunView struct {
	Inventories []ScsiLunInventoryView `json:"inventories,omitempty"`
}

// DetachScsiLunFromHostEventView DetachScsiLunFromHostEvent
type DetachScsiLunFromHostEventView struct {
	Inventory ScsiLunInventoryView `json:"inventory,omitempty"`
}

// AttachScsiLunToVmInstanceEventView AttachScsiLunToVmInstanceEvent
type AttachScsiLunToVmInstanceEventView struct {
	Inventory ScsiLunInventoryView `json:"inventory,omitempty"`
}

// DetachScsiLunFromVmInstanceEventView DetachScsiLunFromVmInstanceEvent
type DetachScsiLunFromVmInstanceEventView struct {
	Inventory ScsiLunInventoryView `json:"inventory,omitempty"`
}

