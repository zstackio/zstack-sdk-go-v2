// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageGroupInventoryView ImageGroup
type ImageGroupInventoryView struct {
	ImageCount *int `json:"imageCount,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

// CreateImageGroupFromVmInstanceEventView CreateImageGroupFromVmInstanceEvent
type CreateImageGroupFromVmInstanceEventView struct {
	Inventory ImageGroupInventoryView `json:"inventory,omitempty"`
}

// CreateImageGroupFromImageEventView CreateImageGroupFromImageEvent
type CreateImageGroupFromImageEventView struct {
	Inventory ImageGroupInventoryView `json:"inventory,omitempty"`
}

// ExpungeImageGroupEventView ExpungeImageGroupEvent
type ExpungeImageGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryImageGroupView QueryImageGroup
type QueryImageGroupView struct {
	Inventories []ImageGroupInventoryView `json:"inventories,omitempty"`
}

// CreateImageGroupFromSnapshotEventView CreateImageGroupFromSnapshotEvent
type CreateImageGroupFromSnapshotEventView struct {
	Inventory ImageGroupInventoryView `json:"inventory,omitempty"`
}

