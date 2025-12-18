// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingRuleAlarmTriggerDetailParam CreateAutoScalingRuleAlarmTrigger detail param
type CreateAutoScalingRuleAlarmTriggerDetailParam struct {
	AlarmUuid string `json:"alarmUuid" validate:"required"`
	TriggerType string `json:"triggerType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	RuleUuid string `json:"ruleUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingRuleAlarmTriggerParam CreateAutoScalingRuleAlarmTrigger request param
type CreateAutoScalingRuleAlarmTriggerParam struct {
	BaseParam
	Params CreateAutoScalingRuleAlarmTriggerDetailParam `json:"params"`
}
