// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingRuleDetailParam DeleteAutoScalingRule detail param
type DeleteAutoScalingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingRuleParam DeleteAutoScalingRule request param
type DeleteAutoScalingRuleParam struct {
	BaseParam
	Params DeleteAutoScalingRuleDetailParam `json:"params"`
}
