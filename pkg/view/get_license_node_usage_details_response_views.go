// Copyright (c) ZStack.io, Inc.

package view

// GetLicenseNodeUsageDetailsView GetLicenseNodeUsageDetails
type GetLicenseNodeUsageDetailsView struct {
	NodeInventory LicenseAuthorizedNodeInventoryView `json:"nodeInventory,omitempty"`
	PlatformLicense LicenseInventoryView `json:"platformLicense,omitempty"`
	AddOns []LicenseAddOnInventoryView `json:"addOns,omitempty"`
	Success bool `json:"success,omitempty"`
}

