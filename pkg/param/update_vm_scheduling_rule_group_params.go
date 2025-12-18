// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmSchedulingRuleGroupDetailParam UpdateVmSchedulingRuleGroup detail param
type UpdateVmSchedulingRuleGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVmSchedulingRuleGroupParam UpdateVmSchedulingRuleGroup request param
type UpdateVmSchedulingRuleGroupParam struct {
	BaseParam
	Params UpdateVmSchedulingRuleGroupDetailParam `json:"params"`
}
