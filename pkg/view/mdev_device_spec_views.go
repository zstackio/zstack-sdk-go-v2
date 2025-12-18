// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MdevDeviceSpecInventoryView MdevDeviceSpec
type MdevDeviceSpecInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Specification string `json:"specification,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	MaxAvailableDevicesPerHost int `json:"maxAvailableDevicesPerHost,omitempty"`
}

