// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LicenseAuthorizedNodeInventoryView LicenseAuthorizedNode
type LicenseAuthorizedNodeInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"appId,omitempty"`
	rest string `json:"ip,omitempty"`
	rest time.Time `json:"lastSyncDate,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

