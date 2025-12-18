// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingGroupAddingNewInstanceRuleDetailParam UpdateAutoScalingGroupAddingNewInstanceRule detail param
type UpdateAutoScalingGroupAddingNewInstanceRuleDetailParam struct {
	AdjustmentType string `json:"adjustmentType,omitempty"`
	AdjustmentValue int `json:"adjustmentValue,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Cooldown int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingGroupAddingNewInstanceRuleParam UpdateAutoScalingGroupAddingNewInstanceRule request param
type UpdateAutoScalingGroupAddingNewInstanceRuleParam struct {
	BaseParam
	Params UpdateAutoScalingGroupAddingNewInstanceRuleDetailParam `json:"params"`
}
