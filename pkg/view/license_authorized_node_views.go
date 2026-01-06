// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LicenseAuthorizedNodeInventoryView LicenseAuthorizedNode
type LicenseAuthorizedNodeInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AppId string `json:"appId,omitempty"`
	Ip string `json:"ip,omitempty"`
	LastSyncDate ZStackTime `json:"lastSyncDate,omitempty"`
	Status string `json:"status,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryLicenseAuthorizedNodeView QueryLicenseAuthorizedNode
type QueryLicenseAuthorizedNodeView struct {
	Inventories []LicenseAuthorizedNodeInventoryView `json:"inventories,omitempty"`
}

