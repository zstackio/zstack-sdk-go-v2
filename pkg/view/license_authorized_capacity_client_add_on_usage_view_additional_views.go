// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAuthorizedCapacityClientAddOnUsageViewView LicenseAuthorizedCapacityClientAddOnUsageView
type LicenseAuthorizedCapacityClientAddOnUsageViewView struct {
	Module *string `json:"module,omitempty"`
	Used int64 `json:"used,omitempty"`
	UsageDetails []LicenseAuthorizedCapacityUsageDetailViewView `json:"usageDetails,omitempty"`
}

