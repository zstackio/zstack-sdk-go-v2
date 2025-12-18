// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingGroupAddingNewInstanceRuleDetailParam UpdateAutoScalingGroupAddingNewInstanceRule详细参数
type UpdateAutoScalingGroupAddingNewInstanceRuleDetailParam struct {
	rest string `json:"adjustmentType,omitempty"`
	rest int `json:"adjustmentValue,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingGroupAddingNewInstanceRuleParam UpdateAutoScalingGroupAddingNewInstanceRule请求参数
type UpdateAutoScalingGroupAddingNewInstanceRuleParam struct {
	BaseParam
	Params UpdateAutoScalingGroupAddingNewInstanceRuleDetailParam `json:"params"` // 详细参数
}

