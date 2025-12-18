// Copyright (c) ZStack.io, Inc.

package view

import "time"

// GuestVmScriptExecutedRecordInventoryView GuestVmScriptExecutedRecord
type GuestVmScriptExecutedRecordInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"scriptUuid,omitempty"`
	rest string `json:"recordName,omitempty"`
	rest int `json:"scriptTimeout,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"executor,omitempty"`
	rest int `json:"executionCount,omitempty"`
	rest int `json:"version,omitempty"`
	rest string `json:"encodingType,omitempty"`
	rest string `json:"scriptContent,omitempty"`
	rest string `json:"renderParams,omitempty"`
	rest time.Time `json:"startTime,omitempty"`
	rest time.Time `json:"endTime,omitempty"`
}

