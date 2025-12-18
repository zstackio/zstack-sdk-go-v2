// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SdnControllerInventoryView SdnController
type SdnControllerInventoryView struct {
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
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

