// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingRuleTriggerDetailParam DeleteAutoScalingRuleTrigger详细参数
type DeleteAutoScalingRuleTriggerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingRuleTriggerParam DeleteAutoScalingRuleTrigger请求参数
type DeleteAutoScalingRuleTriggerParam struct {
	BaseParam
	Params DeleteAutoScalingRuleTriggerDetailParam `json:"params"` // 详细参数
}

