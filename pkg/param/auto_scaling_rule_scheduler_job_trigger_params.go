// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateAutoScalingRuleSchedulerJobTriggerParamDetail CreateAutoScalingRuleSchedulerJobTrigger detail param
type CreateAutoScalingRuleSchedulerJobTriggerParamDetail struct {
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
	Params CreateAutoScalingRuleSchedulerJobTriggerParamDetail `json:"params"`
}
