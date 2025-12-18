// Copyright (c) ZStack.io, Inc.

package view

// GetLatestGuestToolsForVmView GetLatestGuestToolsForVm
type GetLatestGuestToolsForVmView struct {
	Inventory GuestToolsInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

