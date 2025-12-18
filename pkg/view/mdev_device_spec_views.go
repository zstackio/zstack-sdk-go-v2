// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MdevDeviceSpecInventoryView MdevDeviceSpec
type MdevDeviceSpecInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"specification,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int `json:"maxAvailableDevicesPerHost,omitempty"`
}

