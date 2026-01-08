// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OvnControllerInventoryView OvnController
type OvnControllerInventoryView struct {
	BaseInfoView
	BaseTimeView
	RemoteOvn     bool                                `json:"remoteOvn,omitempty"`
	VendorType    string                              `json:"vendorType,omitempty"`
	VendorVersion string                              `json:"vendorVersion,omitempty"`
	Ip            string                              `json:"ip,omitempty"`
	Username      string                              `json:"username,omitempty"`
	Password      string                              `json:"password,omitempty"`
	Status        string                              `json:"status,omitempty"`
	HostRefs      []SdnControllerHostRefInventoryView `json:"hostRefs,omitempty"`
}

// QueryOvnControllerView QueryOvnController
type QueryOvnControllerView struct {
	Inventories []OvnControllerInventoryView `json:"inventories,omitempty"`
}
