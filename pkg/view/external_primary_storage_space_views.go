// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ExternalPrimaryStorageSpaceInventoryView ExternalPrimaryStorageSpace
type ExternalPrimaryStorageSpaceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	LocationUrl string `json:"locationUrl,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

