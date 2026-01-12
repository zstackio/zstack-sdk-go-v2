// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsSecurityGroupRuleInventoryView EcsSecurityGroupRule
type EcsSecurityGroupRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	EcsSecurityGroupUuid *string `json:"ecsSecurityGroupUuid,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	PortRange *string `json:"portRange,omitempty"`
	CidrIp *string `json:"cidrIp,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Direction *string `json:"direction,omitempty"`
	NicType *string `json:"nicType,omitempty"`
	Policy *string `json:"policy,omitempty"`
	Description *string `json:"description,omitempty"`
}

// SyncEcsSecurityGroupRuleFromRemoteEventView SyncEcsSecurityGroupRuleFromRemoteEvent
type SyncEcsSecurityGroupRuleFromRemoteEventView struct {
	Inventories []EcsSecurityGroupRuleInventoryView `json:"inventories,omitempty"`
}

// QueryEcsSecurityGroupRuleFromLocalView QueryEcsSecurityGroupRuleFromLocal
type QueryEcsSecurityGroupRuleFromLocalView struct {
	Inventories []EcsSecurityGroupRuleInventoryView `json:"inventories,omitempty"`
}

// CreateEcsSecurityGroupRuleRemoteEventView CreateEcsSecurityGroupRuleRemoteEvent
type CreateEcsSecurityGroupRuleRemoteEventView struct {
	Inventory EcsSecurityGroupRuleInventoryView `json:"inventory,omitempty"`
}

