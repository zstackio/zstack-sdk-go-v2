// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RootVolumeUsageInventoryView RootVolumeUsage
type RootVolumeUsageInventoryView struct {
	Id int64 `json:"id,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VolumeStatus string `json:"volumeStatus,omitempty"`
	VolumeName int64 `json:"volumeName,omitempty"`
	VolumeSize int64 `json:"volumeSize,omitempty"`
	Inventory string `json:"inventory,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

