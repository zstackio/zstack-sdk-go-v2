// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAuthorizedCapacityInventoryView LicenseAuthorizedCapacity
type LicenseAuthorizedCapacityInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id *int64 `json:"id,omitempty"`
	NodeUuid *string `json:"nodeUuid,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	ResourceInfo *string `json:"resourceInfo,omitempty"`
	QuotaType *string `json:"quotaType,omitempty"`
	Quota *int64 `json:"quota,omitempty"`
	LicenseType *string `json:"licenseType,omitempty"`
	Type *string `json:"type,omitempty"`
}

// GetLicenseAuthorizedCapacityView GetLicenseAuthorizedCapacity
type GetLicenseAuthorizedCapacityView struct {
	Total TotalLicenseAuthorizedCapacityViewView `json:"total,omitempty"`
	Clients []LicenseAuthorizedCapacityClientUsageViewView `json:"clients,omitempty"`
	Server LicenseAuthorizedCapacityServerUsageViewView `json:"server,omitempty"`
	Success bool `json:"success,omitempty"`
}

