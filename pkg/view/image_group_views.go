// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ImageGroupInventoryView ImageGroup
type ImageGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	ImageCount int `json:"imageCount,omitempty"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
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

