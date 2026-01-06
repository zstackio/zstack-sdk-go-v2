// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MttyDeviceInventoryView MttyDevice
type MttyDeviceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	VirtStatus string `json:"virtStatus,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryMttyDeviceView QueryMttyDevice
type QueryMttyDeviceView struct {
	Inventories []MttyDeviceInventoryView `json:"inventories,omitempty"`
}

