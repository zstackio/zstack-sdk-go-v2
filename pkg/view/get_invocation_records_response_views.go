// Copyright (c) ZStack.io, Inc.

package view

// GetInvocationRecordsView GetInvocationRecords
type GetInvocationRecordsView struct {
	Inventories []InvocationRecordView `json:"inventories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

