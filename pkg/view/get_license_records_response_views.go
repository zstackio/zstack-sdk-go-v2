// Copyright (c) ZStack.io, Inc.

package view

// GetLicenseRecordsView GetLicenseRecords
type GetLicenseRecordsView struct {
	Inventories []LicenseInventoryView `json:"inventories,omitempty"`
	Total int64 `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

