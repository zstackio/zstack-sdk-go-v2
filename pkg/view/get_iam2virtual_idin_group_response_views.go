// Copyright (c) ZStack.io, Inc.

package view

// GetIAM2VirtualIDInGroupView GetIAM2VirtualIDInGroup
type GetIAM2VirtualIDInGroupView struct {
	Inventories []IAM2VirtualIDInventoryView `json:"inventories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

