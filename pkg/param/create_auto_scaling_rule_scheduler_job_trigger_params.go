// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingRuleSchedulerJobTriggerDetailParam CreateAutoScalingRuleSchedulerJobTrigger detail param
type CreateAutoScalingRuleSchedulerJobTriggerDetailParam struct {
	SchedulerJobUuid string `json:"schedulerJobUuid" validate:"required"`
	TriggerType string `json:"triggerType,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	RuleUuid string `json:"ruleUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingRuleSchedulerJobTriggerParam CreateAutoScalingRuleSchedulerJobTrigger request param
type CreateAutoScalingRuleSchedulerJobTriggerParam struct {
	BaseParam
	Params CreateAutoScalingRuleSchedulerJobTriggerDetailParam `json:"params"`
}
