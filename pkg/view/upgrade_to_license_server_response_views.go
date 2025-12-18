// Copyright (c) ZStack.io, Inc.

package view

// UpgradeToLicenseServerEventView UpgradeToLicenseServerEvent
type UpgradeToLicenseServerEventView struct {
	Inventory LicenseAuthorizedNodeInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

