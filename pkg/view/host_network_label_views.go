// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostNetworkLabelInventoryView HostNetworkLabel
type HostNetworkLabelInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
	System bool `json:"system,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

