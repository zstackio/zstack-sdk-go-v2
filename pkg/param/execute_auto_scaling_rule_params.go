// Copyright (c) ZStack.io, Inc.

package param

// ExecuteAutoScalingRuleDetailParam ExecuteAutoScalingRule detail param
type ExecuteAutoScalingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExecuteAutoScalingRuleParam ExecuteAutoScalingRule request param
type ExecuteAutoScalingRuleParam struct {
	BaseParam
	Params ExecuteAutoScalingRuleDetailParam `json:"params"`
}
