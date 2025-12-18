// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EcsSecurityGroupRuleInventoryView EcsSecurityGroupRule
type EcsSecurityGroupRuleInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"ecsSecurityGroupUuid,omitempty"`
	rest string `json:"protocol,omitempty"`
	rest string `json:"portRange,omitempty"`
	rest string `json:"cidrIp,omitempty"`
	rest string `json:"priority,omitempty"`
	rest string `json:"direction,omitempty"`
	rest string `json:"nicType,omitempty"`
	rest string `json:"policy,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

