// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostNetworkLabelInventoryView HostNetworkLabel
type HostNetworkLabelInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"serviceType,omitempty"`
	rest bool `json:"system,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

