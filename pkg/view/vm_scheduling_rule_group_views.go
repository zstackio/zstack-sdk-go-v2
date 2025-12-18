// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmSchedulingRuleGroupInventoryView VmSchedulingRuleGroup
type VmSchedulingRuleGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"appliance,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

