// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SdnControllerInventoryView SdnController
type SdnControllerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VendorType *string `json:"vendorType,omitempty"`
	VendorVersion *string `json:"vendorVersion,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Status string `json:"status,omitempty"`
	HostRefs []SdnControllerHostRefInventoryView `json:"hostRefs,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// RemoveSdnControllerEventView RemoveSdnControllerEvent
type RemoveSdnControllerEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddSdnControllerEventView AddSdnControllerEvent
type AddSdnControllerEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// UpdateSdnControllerEventView UpdateSdnControllerEvent
type UpdateSdnControllerEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// ChangeSdnControllerEventView ChangeSdnControllerEvent
type ChangeSdnControllerEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// ReconnectSdnControllerEventView ReconnectSdnControllerEvent
type ReconnectSdnControllerEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// QuerySdnControllerView QuerySdnController
type QuerySdnControllerView struct {
	Inventories []SdnControllerInventoryView `json:"inventories,omitempty"`
}

