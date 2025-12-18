// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingGroupRemovalInstanceRuleDetailParam UpdateAutoScalingGroupRemovalInstanceRule详细参数
type UpdateAutoScalingGroupRemovalInstanceRuleDetailParam struct {
	rest string `json:"adjustmentType,omitempty"`
	rest int `json:"adjustmentValue,omitempty"`
	rest string `json:"removalPolicy,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingGroupRemovalInstanceRuleParam UpdateAutoScalingGroupRemovalInstanceRule请求参数
type UpdateAutoScalingGroupRemovalInstanceRuleParam struct {
	BaseParam
	Params UpdateAutoScalingGroupRemovalInstanceRuleDetailParam `json:"params"` // 详细参数
}

