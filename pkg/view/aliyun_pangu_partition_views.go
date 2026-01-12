// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunPanguPartitionInventoryView AliyunPanguPartition
type AliyunPanguPartitionInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountUuid *string `json:"accountUuid,omitempty"`
	IdentityZoneUuid *string `json:"identityZoneUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	AppName *string `json:"appName,omitempty"`
	PartitionName *string `json:"partitionName,omitempty"`
}

// AddAliyunPanguPartitionEventView AddAliyunPanguPartitionEvent
type AddAliyunPanguPartitionEventView struct {
	Inventory AliyunPanguPartitionInventoryView `json:"inventory,omitempty"`
}

// DeleteAliyunPanguPartitionEventView DeleteAliyunPanguPartitionEvent
type DeleteAliyunPanguPartitionEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryAliyunPanguPartitionView QueryAliyunPanguPartition
type QueryAliyunPanguPartitionView struct {
	Inventories []AliyunPanguPartitionInventoryView `json:"inventories,omitempty"`
}

// UpdateAliyunPanguPartitionEventView UpdateAliyunPanguPartitionEvent
type UpdateAliyunPanguPartitionEventView struct {
	Inventory AliyunPanguPartitionInventoryView `json:"inventory,omitempty"`
}

