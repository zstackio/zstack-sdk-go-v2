// Copyright (c) ZStack.io, Inc.

package param

// ExecuteAutoScalingRuleDetailParam ExecuteAutoScalingRule详细参数
type ExecuteAutoScalingRuleDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ExecuteAutoScalingRuleParam ExecuteAutoScalingRule请求参数
type ExecuteAutoScalingRuleParam struct {
	BaseParam
	Params ExecuteAutoScalingRuleDetailParam `json:"params"` // 详细参数
}

