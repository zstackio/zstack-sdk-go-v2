// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// TotalLicenseAuthorizedCapacityViewView TotalLicenseAuthorizedCapacityView
type TotalLicenseAuthorizedCapacityViewView struct {
	ServerAppId *string `json:"serverAppId,omitempty"`
	ServerAuthorizedNodeUuid *string `json:"serverAuthorizedNodeUuid,omitempty"`
	ServerInventory LicenseAuthorizedNodeInventoryView `json:"serverInventory,omitempty"`
	PlatformLicense LicenseInventoryView `json:"platformLicense,omitempty"`
	AddOns []LicenseAddOnInventoryView `json:"addOns,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

