// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OvnControllerInventoryView OvnController
type OvnControllerInventoryView struct {
	RemoteOvn bool `json:"remoteOvn,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	VendorType string `json:"vendorType,omitempty"`
	VendorVersion string `json:"vendorVersion,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Ip string `json:"ip,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Status string `json:"status,omitempty"`
	HostRefs []SdnControllerHostRefInventoryView `json:"hostRefs,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryOvnControllerView QueryOvnController
type QueryOvnControllerView struct {
	Inventories []OvnControllerInventoryView `json:"inventories,omitempty"`
}

