// Copyright (c) ZStack.io, Inc.

package param

// DeleteAutoScalingRuleTriggerDetailParam DeleteAutoScalingRuleTrigger detail param
type DeleteAutoScalingRuleTriggerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingRuleTriggerParam DeleteAutoScalingRuleTrigger request param
type DeleteAutoScalingRuleTriggerParam struct {
	BaseParam
	Params DeleteAutoScalingRuleTriggerDetailParam `json:"params"`
}
