// Copyright (c) ZStack.io, Inc.

package view

// RegisterLicenseServerEventView RegisterLicenseServerEvent
type RegisterLicenseServerEventView struct {
	LicenseClient LicenseAuthorizedNodeInventoryView `json:"licenseClient,omitempty"`
	LicenseServer LicenseAuthorizedNodeInventoryView `json:"licenseServer,omitempty"`
	Success bool `json:"success,omitempty"`
}

