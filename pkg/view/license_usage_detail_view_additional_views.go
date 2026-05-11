// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LicenseUsageDetailViewView LicenseUsageDetailView
type LicenseUsageDetailViewView struct {
	UsedBy string `json:"usedBy,omitempty"`
	Usage int64 `json:"usage,omitempty"`
	ResourceInfo string `json:"resourceInfo,omitempty"`
}

