// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AffinityGroupInventoryView AffinityGroup
type AffinityGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"policy,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"appliance,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []AffinityGroupUsageInventoryView `json:"usages,omitempty"`
}

