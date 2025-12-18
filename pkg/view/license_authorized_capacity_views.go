// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LicenseAuthorizedCapacityInventoryView LicenseAuthorizedCapacity
type LicenseAuthorizedCapacityInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"nodeUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"quotaType,omitempty"`
	rest int64 `json:"quota,omitempty"`
	rest string `json:"licenseType,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

