// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsImageUsageInventoryView EcsImageUsage
type EcsImageUsageInventoryView struct {
	Id int `json:"id,omitempty"`
	EcsImageUuid string `json:"ecsImageUuid,omitempty"`
	SnapshotUuidOfCreatedImage string `json:"snapshotUuidOfCreatedImage,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

