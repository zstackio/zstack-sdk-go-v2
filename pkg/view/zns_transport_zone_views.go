// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZnsTransportZoneInventoryView ZnsTransportZone
type ZnsTransportZoneInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	PhysicalNetwork string `json:"physicalNetwork,omitempty"`
	Status string `json:"status,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Tags string `json:"tags,omitempty"`
	ZnsSdnControllerUuid string `json:"znsSdnControllerUuid,omitempty"`
}

