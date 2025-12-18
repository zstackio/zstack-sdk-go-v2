// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyRouteRuleSetVRouterRefInventoryView PolicyRouteRuleSetVRouterRef
type PolicyRouteRuleSetVRouterRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"vRouterUuid,omitempty"`
	rest string `json:"ruleSetUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

