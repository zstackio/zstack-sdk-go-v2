// Copyright (c) ZStack.io, Inc.

package view

// GetLicenseInfoView GetLicenseInfo
type GetLicenseInfoView struct {
	Inventory LicenseInventoryView `json:"inventory,omitempty"`
	Additions []interface{} `json:"additions,omitempty"`
}

