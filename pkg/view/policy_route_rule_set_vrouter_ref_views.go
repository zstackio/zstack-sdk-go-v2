// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteRuleSetVRouterRefInventoryView PolicyRouteRuleSetVRouterRef
type PolicyRouteRuleSetVRouterRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryPolicyRouteRuleSetVRouterRefView QueryPolicyRouteRuleSetVRouterRef
type QueryPolicyRouteRuleSetVRouterRefView struct {
	Inventories []PolicyRouteRuleSetVRouterRefInventoryView `json:"inventories,omitempty"`
}

