// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmSchedulingRuleGroupRefInventoryView VmSchedulingRuleGroupRef
type VmSchedulingRuleGroupRefInventoryView struct {
	rest string `json:"vmGroupUuid,omitempty"`
	rest string `json:"vmUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

