// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingRuleDetailParam UpdateAutoScalingRule detail param
type UpdateAutoScalingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Cooldown int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingRuleParam UpdateAutoScalingRule request param
type UpdateAutoScalingRuleParam struct {
	BaseParam
	Params UpdateAutoScalingRuleDetailParam `json:"params"`
}
