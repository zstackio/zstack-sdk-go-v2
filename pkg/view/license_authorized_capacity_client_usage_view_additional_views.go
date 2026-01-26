// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAuthorizedCapacityClientUsageViewView LicenseAuthorizedCapacityClientUsageView
type LicenseAuthorizedCapacityClientUsageViewView struct {
	ClientAppId string `json:"clientAppId,omitempty"`
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid,omitempty"`
	ClientInventory LicenseAuthorizedNodeInventoryView `json:"clientInventory,omitempty"`
	PlatformUsed int64 `json:"platformUsed,omitempty"`
	PlatformUsageDetails []LicenseAuthorizedCapacityUsageDetailViewView `json:"platformUsageDetails,omitempty"`
	AddOns []LicenseAuthorizedCapacityClientAddOnUsageViewView `json:"addOns,omitempty"`
}

