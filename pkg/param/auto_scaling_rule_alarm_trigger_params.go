// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingRuleAlarmTriggerDetailParam CreateAutoScalingRuleAlarmTrigger详细参数
type CreateAutoScalingRuleAlarmTriggerDetailParam struct {
	rest string `json:"alarmUuid" validate:"required"` // 必填
	rest string `json:"triggerType,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"ruleUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingRuleAlarmTriggerParam CreateAutoScalingRuleAlarmTrigger请求参数
type CreateAutoScalingRuleAlarmTriggerParam struct {
	BaseParam
	Params CreateAutoScalingRuleAlarmTriggerDetailParam `json:"params"` // 详细参数
}

