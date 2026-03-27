// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// InvocationRecordView InvocationRecord
type InvocationRecordView struct {
	StartTime time.Time `json:"startTime,omitempty"`
	Status string `json:"status,omitempty"`
	ScriptUuid string `json:"scriptUuid,omitempty"`
	ScriptType string `json:"scriptType,omitempty"`
	ScriptContent string `json:"scriptContent,omitempty"`
	ScriptName string `json:"scriptName,omitempty"`
	Description string `json:"description,omitempty"`
	RecordUuid string `json:"recordUuid,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	Inventories []InvocationRecordDetailView `json:"inventories,omitempty"`
}

