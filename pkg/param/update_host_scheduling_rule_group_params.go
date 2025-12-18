// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostSchedulingRuleGroupDetailParam UpdateHostSchedulingRuleGroup detail param
type UpdateHostSchedulingRuleGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateHostSchedulingRuleGroupParam UpdateHostSchedulingRuleGroup request param
type UpdateHostSchedulingRuleGroupParam struct {
	BaseParam
	Params UpdateHostSchedulingRuleGroupDetailParam `json:"params"`
}
