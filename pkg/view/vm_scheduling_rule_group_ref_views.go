// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmSchedulingRuleGroupRefInventoryView VmSchedulingRuleGroupRef
type VmSchedulingRuleGroupRefInventoryView struct {
	VmGroupUuid *string `json:"vmGroupUuid,omitempty"`
	VmUuid *string `json:"vmUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

