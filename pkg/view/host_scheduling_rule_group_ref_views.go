// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostSchedulingRuleGroupRefInventoryView HostSchedulingRuleGroupRef
type HostSchedulingRuleGroupRefInventoryView struct {
	HostGroupUuid string `json:"hostGroupUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

