// Copyright (c) ZStack.io, Inc.

package view

import "time"

// GuestVmScriptExecutedRecordDetailInventoryView GuestVmScriptExecutedRecordDetail
type GuestVmScriptExecutedRecordDetailInventoryView struct {
	rest string `json:"recordUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest string `json:"vmName,omitempty"`
	rest string `json:"status,omitempty"`
	rest int `json:"exitCode,omitempty"`
	rest string `json:"stdout,omitempty"`
	rest string `json:"errCause,omitempty"`
	rest string `json:"stderr,omitempty"`
	rest time.Time `json:"startTime,omitempty"`
	rest time.Time `json:"endTime,omitempty"`
}

