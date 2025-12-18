// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyRouteRuleSetL3RefInventoryView PolicyRouteRuleSetL3Ref
type PolicyRouteRuleSetL3RefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"ruleSetUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

