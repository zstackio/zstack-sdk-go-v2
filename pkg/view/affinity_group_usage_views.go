// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AffinityGroupUsageInventoryView AffinityGroupUsage
type AffinityGroupUsageInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AffinityGroupUuid string `json:"affinityGroupUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

