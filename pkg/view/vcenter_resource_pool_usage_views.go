// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VCenterResourcePoolUsageInventoryView VCenterResourcePoolUsage
type VCenterResourcePoolUsageInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VCenterResourcePoolUuid string `json:"vCenterResourcePoolUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

