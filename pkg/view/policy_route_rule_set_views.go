// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyRouteRuleSetInventoryView PolicyRouteRuleSet
type PolicyRouteRuleSetInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []PolicyRouteRuleInventoryView `json:"rules,omitempty"`
	rest []PolicyRouteRuleSetL3RefInventoryView `json:"l3Refs,omitempty"`
}

