// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmSchedulingRuleDetailParam UpdateVmSchedulingRule detail param
type UpdateVmSchedulingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Mode string `json:"mode,omitempty"`
}

// UpdateVmSchedulingRuleParam UpdateVmSchedulingRule request param
type UpdateVmSchedulingRuleParam struct {
	BaseParam
	Params UpdateVmSchedulingRuleDetailParam `json:"params"`
}
