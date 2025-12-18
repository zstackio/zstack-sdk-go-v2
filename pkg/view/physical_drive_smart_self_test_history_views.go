// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PhysicalDriveSmartSelfTestHistoryInventoryView PhysicalDriveSmartSelfTestHistory
type PhysicalDriveSmartSelfTestHistoryInventoryView struct {
	Id int64 `json:"id,omitempty"`
	RaidPhysicalDriveUuid string `json:"raidPhysicalDriveUuid,omitempty"`
	RunningState string `json:"runningState,omitempty"`
	TestResult string `json:"testResult,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

