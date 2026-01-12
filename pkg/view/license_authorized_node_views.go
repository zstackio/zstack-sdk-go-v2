// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAuthorizedNodeInventoryView LicenseAuthorizedNode
type LicenseAuthorizedNodeInventoryView struct {
	BaseInfoView
	BaseTimeView
	AppId *string `json:"appId,omitempty"`
	Ip *string `json:"ip,omitempty"`
	LastSyncDate *time.Time `json:"lastSyncDate,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
}

// QueryLicenseAuthorizedNodeView QueryLicenseAuthorizedNode
type QueryLicenseAuthorizedNodeView struct {
	Inventories []LicenseAuthorizedNodeInventoryView `json:"inventories,omitempty"`
}

