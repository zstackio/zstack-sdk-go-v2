// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAppIdRefInventoryView LicenseAppIdRef
type LicenseAppIdRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	LicenseId string `json:"licenseId,omitempty"`
	AppId string `json:"appId,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

