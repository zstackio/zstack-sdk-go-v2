// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAuthorizedCapacityUsageDetailViewView LicenseAuthorizedCapacityUsageDetailView
type LicenseAuthorizedCapacityUsageDetailViewView struct {
	ResourceUuid string `json:"resourceUuid,omitempty"`
	Used int64 `json:"used,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

