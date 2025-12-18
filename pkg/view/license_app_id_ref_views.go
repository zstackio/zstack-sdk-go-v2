// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LicenseAppIdRefInventoryView LicenseAppIdRef
type LicenseAppIdRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"licenseId,omitempty"`
	rest string `json:"appId,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

