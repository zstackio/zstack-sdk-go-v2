// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostSchedulingRuleGroupInventoryView HostSchedulingRuleGroup
type HostSchedulingRuleGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

