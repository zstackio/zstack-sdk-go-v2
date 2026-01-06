// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HbaDeviceInventoryView HbaDevice
type HbaDeviceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	HbaType string `json:"hbaType,omitempty"`
	CreateDate string `json:"createDate,omitempty"`
	LastOpDate string `json:"lastOpDate,omitempty"`
}

// QueryFcHbaDeviceView QueryFcHbaDevice
type QueryFcHbaDeviceView struct {
	Inventories []HbaDeviceInventoryView `json:"inventories,omitempty"`
}

