// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EcsImageUsageInventoryView EcsImageUsage
type EcsImageUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int `json:"id,omitempty"`
	EcsImageUuid string `json:"ecsImageUuid,omitempty"`
	SnapshotUuidOfCreatedImage string `json:"snapshotUuidOfCreatedImage,omitempty"`
}

