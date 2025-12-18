// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AffinityGroupUsageInventoryView AffinityGroupUsage
type AffinityGroupUsageInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"affinityGroupUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

