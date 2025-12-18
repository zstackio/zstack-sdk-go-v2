// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmSchedulingRuleRefInventoryView VmSchedulingRuleRef
type VmSchedulingRuleRefInventoryView struct {
	rest string `json:"vmGroupUuid,omitempty"`
	rest string `json:"hostGroupUuid,omitempty"`
	rest string `json:"vmSchedulingRuleUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

