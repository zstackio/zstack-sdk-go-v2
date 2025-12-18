// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VCenterResourcePoolUsageInventoryView VCenterResourcePoolUsage
type VCenterResourcePoolUsageInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vCenterResourcePoolUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

