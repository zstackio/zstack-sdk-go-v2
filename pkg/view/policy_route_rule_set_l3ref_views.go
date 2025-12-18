// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteRuleSetL3RefInventoryView PolicyRouteRuleSetL3Ref
type PolicyRouteRuleSetL3RefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

