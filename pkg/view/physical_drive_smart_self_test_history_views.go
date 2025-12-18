// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PhysicalDriveSmartSelfTestHistoryInventoryView PhysicalDriveSmartSelfTestHistory
type PhysicalDriveSmartSelfTestHistoryInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"raidPhysicalDriveUuid,omitempty"`
	rest string `json:"runningState,omitempty"`
	rest string `json:"testResult,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

