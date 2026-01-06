// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmSchedulingRuleRefInventoryView VmSchedulingRuleRef
type VmSchedulingRuleRefInventoryView struct {
	VmGroupUuid string `json:"vmGroupUuid,omitempty"`
	HostGroupUuid string `json:"hostGroupUuid,omitempty"`
	VmSchedulingRuleUuid string `json:"vmSchedulingRuleUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

