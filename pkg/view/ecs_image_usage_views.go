// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsImageUsageInventoryView EcsImageUsage
type EcsImageUsageInventoryView struct {
	Id int `json:"id,omitempty"`
	EcsImageUuid string `json:"ecsImageUuid,omitempty"`
	SnapshotUuidOfCreatedImage string `json:"snapshotUuidOfCreatedImage,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

