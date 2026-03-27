// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BareMetal2ChassisDiskInventoryView BareMetal2ChassisDisk
type BareMetal2ChassisDiskInventoryView struct {
	BaseInfoView
	BaseTimeView
	ChassisUuid string `json:"chassisUuid,omitempty"`
	DiskSize int64 `json:"diskSize,omitempty"`
	Type string `json:"type,omitempty"`
	Wwn string `json:"wwn,omitempty"`
}

