// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GuestVmScriptExecutedRecordDetailInventoryView GuestVmScriptExecutedRecordDetail
type GuestVmScriptExecutedRecordDetailInventoryView struct {
	RecordUuid string `json:"recordUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VmName string `json:"vmName,omitempty"`
	Status string `json:"status,omitempty"`
	ExitCode int `json:"exitCode,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	ErrCause string `json:"errCause,omitempty"`
	Stderr string `json:"stderr,omitempty"`
	StartTime ZStackTime `json:"startTime,omitempty"`
	EndTime ZStackTime `json:"endTime,omitempty"`
}

// QueryGuestVmScriptExecutedRecordDetailView QueryGuestVmScriptExecutedRecordDetail
type QueryGuestVmScriptExecutedRecordDetailView struct {
	Inventories []GuestVmScriptExecutedRecordDetailInventoryView `json:"inventories,omitempty"`
}

