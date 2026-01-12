// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunDiskInventoryView AliyunDisk
type AliyunDiskInventoryView struct {
	BaseInfoView
	BaseTimeView
	DiskId *string `json:"diskId,omitempty"`
	Description *string `json:"description,omitempty"`
	IdentityZoneUuid *string `json:"identityZoneUuid,omitempty"`
	EcsInstanceUuid *string `json:"ecsInstanceUuid,omitempty"`
	DiskCategory *string `json:"diskCategory,omitempty"`
	DiskType *string `json:"diskType,omitempty"`
	DiskChargeType *string `json:"diskChargeType,omitempty"`
	Status *string `json:"status,omitempty"`
	SizeWithGB *int `json:"sizeWithGB,omitempty"`
	DeviceInfo *string `json:"deviceInfo,omitempty"`
}

// QueryAliyunDiskFromLocalView QueryAliyunDiskFromLocal
type QueryAliyunDiskFromLocalView struct {
	Inventories []AliyunDiskInventoryView `json:"inventories,omitempty"`
}

// UpdateAliyunDiskEventView UpdateAliyunDiskEvent
type UpdateAliyunDiskEventView struct {
	Inventory AliyunDiskInventoryView `json:"inventory,omitempty"`
}

// CreateAliyunDiskFromRemoteEventView CreateAliyunDiskFromRemoteEvent
type CreateAliyunDiskFromRemoteEventView struct {
	Inventory AliyunDiskInventoryView `json:"inventory,omitempty"`
}

// AttachAliyunDiskToEcsEventView AttachAliyunDiskToEcsEvent
type AttachAliyunDiskToEcsEventView struct {
	Inventory AliyunDiskInventoryView `json:"inventory,omitempty"`
}

// SyncDiskFromAliyunFromRemoteEventView SyncDiskFromAliyunFromRemoteEvent
type SyncDiskFromAliyunFromRemoteEventView struct {
	Inventories []AliyunDiskInventoryView `json:"inventories,omitempty"`
}

