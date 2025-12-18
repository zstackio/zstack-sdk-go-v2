// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingGroupRemovalInstanceRuleDetailParam UpdateAutoScalingGroupRemovalInstanceRule detail param
type UpdateAutoScalingGroupRemovalInstanceRuleDetailParam struct {
	AdjustmentType string `json:"adjustmentType,omitempty"`
	AdjustmentValue int `json:"adjustmentValue,omitempty"`
	RemovalPolicy string `json:"removalPolicy,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Cooldown int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingGroupRemovalInstanceRuleParam UpdateAutoScalingGroupRemovalInstanceRule request param
type UpdateAutoScalingGroupRemovalInstanceRuleParam struct {
	BaseParam
	Params UpdateAutoScalingGroupRemovalInstanceRuleDetailParam `json:"params"`
}
