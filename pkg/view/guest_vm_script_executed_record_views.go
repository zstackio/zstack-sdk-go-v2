// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GuestVmScriptExecutedRecordInventoryView GuestVmScriptExecutedRecord
type GuestVmScriptExecutedRecordInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ScriptUuid string `json:"scriptUuid,omitempty"`
	RecordName string `json:"recordName,omitempty"`
	ScriptTimeout int `json:"scriptTimeout,omitempty"`
	Status string `json:"status,omitempty"`
	Executor string `json:"executor,omitempty"`
	ExecutionCount int `json:"executionCount,omitempty"`
	Version int `json:"version,omitempty"`
	EncodingType string `json:"encodingType,omitempty"`
	ScriptContent string `json:"scriptContent,omitempty"`
	RenderParams string `json:"renderParams,omitempty"`
	StartTime ZStackTime `json:"startTime,omitempty"`
	EndTime ZStackTime `json:"endTime,omitempty"`
}

// QueryGuestVmScriptExecutedRecordView QueryGuestVmScriptExecutedRecord
type QueryGuestVmScriptExecutedRecordView struct {
	Inventories []GuestVmScriptExecutedRecordInventoryView `json:"inventories,omitempty"`
}

