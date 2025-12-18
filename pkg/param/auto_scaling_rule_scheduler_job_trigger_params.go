// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingRuleSchedulerJobTriggerDetailParam CreateAutoScalingRuleSchedulerJobTrigger详细参数
type CreateAutoScalingRuleSchedulerJobTriggerDetailParam struct {
	rest string `json:"schedulerJobUuid" validate:"required"` // 必填
	rest string `json:"triggerType,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"ruleUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingRuleSchedulerJobTriggerParam CreateAutoScalingRuleSchedulerJobTrigger请求参数
type CreateAutoScalingRuleSchedulerJobTriggerParam struct {
	BaseParam
	Params CreateAutoScalingRuleSchedulerJobTriggerDetailParam `json:"params"` // 详细参数
}

