// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VCenterBackupStorageInventoryView VCenterBackupStorage
type VCenterBackupStorageInventoryView struct {
	VCenterUuid string `json:"vCenterUuid,omitempty"`
	Datastore string `json:"datastore,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	AttachedZoneUuids []string `json:"attachedZoneUuids,omitempty"`
}

// QueryVCenterBackupStorageView QueryVCenterBackupStorage
type QueryVCenterBackupStorageView struct {
	Inventories []VCenterBackupStorageInventoryView `json:"inventories,omitempty"`
}

