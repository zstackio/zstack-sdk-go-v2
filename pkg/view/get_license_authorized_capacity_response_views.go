// Copyright (c) ZStack.io, Inc.

package view

// GetLicenseAuthorizedCapacityView GetLicenseAuthorizedCapacity
type GetLicenseAuthorizedCapacityView struct {
	Total TotalLicenseAuthorizedCapacityViewView `json:"total,omitempty"`
	Clients []LicenseAuthorizedCapacityClientUsageViewView `json:"clients,omitempty"`
	Server LicenseAuthorizedCapacityServerUsageViewView `json:"server,omitempty"`
	Success bool `json:"success,omitempty"`
}

