// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostSchedulingRuleGroupRefInventoryView HostSchedulingRuleGroupRef
type HostSchedulingRuleGroupRefInventoryView struct {
	rest string `json:"hostGroupUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

