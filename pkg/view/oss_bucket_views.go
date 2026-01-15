// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OssBucketInventoryView OssBucket
type OssBucketInventoryView struct {
	BaseInfoView
	BaseTimeView
	BucketName string `json:"bucketName,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	Current string `json:"current,omitempty"`
	RegionName string `json:"regionName,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateOssBucketRemoteEventView CreateOssBucketRemoteEvent
type CreateOssBucketRemoteEventView struct {
	Inventory OssBucketInventoryView `json:"inventory,omitempty"`
}

// AddOssBucketFromRemoteEventView AddOssBucketFromRemoteEvent
type AddOssBucketFromRemoteEventView struct {
	Inventory OssBucketInventoryView `json:"inventory,omitempty"`
}

// AttachOssBucketToEcsDataCenterEventView AttachOssBucketToEcsDataCenterEvent
type AttachOssBucketToEcsDataCenterEventView struct {
	Inventory OssBucketInventoryView `json:"inventory,omitempty"`
}

// UpdateOssBucketEventView UpdateOssBucketEvent
type UpdateOssBucketEventView struct {
	Inventory OssBucketInventoryView `json:"inventory,omitempty"`
}

// QueryOssBucketFileNameView QueryOssBucketFileName
type QueryOssBucketFileNameView struct {
	Inventories []OssBucketInventoryView `json:"inventories,omitempty"`
}

// DetachOssBucketFromEcsDataCenterEventView DetachOssBucketFromEcsDataCenterEvent
type DetachOssBucketFromEcsDataCenterEventView struct {
	Inventory OssBucketInventoryView `json:"inventory,omitempty"`
}

