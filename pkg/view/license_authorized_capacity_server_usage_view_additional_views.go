// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAuthorizedCapacityServerUsageViewView LicenseAuthorizedCapacityServerUsageView
type LicenseAuthorizedCapacityServerUsageViewView struct {
	PlatformUsed int64 `json:"platformUsed,omitempty"`
	AddOns []LicenseAuthorizedCapacityClientAddOnUsageViewView `json:"addOns,omitempty"`
}

