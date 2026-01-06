// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageStoreBackupStorageInventoryView ImageStoreBackupStorage
type ImageStoreBackupStorageInventoryView struct {
	Hostname string `json:"hostname,omitempty"`
	Username string `json:"username,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
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

// AddImageStoreBackupStorageEventView AddImageStoreBackupStorageEvent
type AddImageStoreBackupStorageEventView struct {
	Inventory ImageStoreBackupStorageInventoryView `json:"inventory,omitempty"`
}

// UpdateBackupStorageEventView UpdateBackupStorageEvent
type UpdateBackupStorageEventView struct {
	Inventory BackupStorageInventoryView `json:"inventory,omitempty"`
}

// QueryImageStoreBackupStorageView QueryImageStoreBackupStorage
type QueryImageStoreBackupStorageView struct {
	Inventories []ImageStoreBackupStorageInventoryView `json:"inventories,omitempty"`
}

// ReconnectImageStoreBackupStorageEventView ReconnectImageStoreBackupStorageEvent
type ReconnectImageStoreBackupStorageEventView struct {
	Inventory ImageStoreBackupStorageInventoryView `json:"inventory,omitempty"`
}

