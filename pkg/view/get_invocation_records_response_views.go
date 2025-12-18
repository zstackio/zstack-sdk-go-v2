// Copyright (c) ZStack.io, Inc.

package view

// GetInvocationRecordsView GetInvocationRecords
type GetInvocationRecordsView struct {
	Inventories []interface{} `json:"inventories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

