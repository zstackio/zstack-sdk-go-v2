// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteRuleSetInventoryView PolicyRouteRuleSet
type PolicyRouteRuleSetInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Rules []PolicyRouteRuleInventoryView `json:"rules,omitempty"`
	L3Refs []PolicyRouteRuleSetL3RefInventoryView `json:"l3Refs,omitempty"`
}

