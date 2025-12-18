// Copyright (c) ZStack.io, Inc.

package param

// AddSchedulerJobGroupToSchedulerTriggerDetailParam AddSchedulerJobGroupToSchedulerTrigger详细参数
type AddSchedulerJobGroupToSchedulerTriggerDetailParam struct {
	rest string `json:"schedulerJobGroupUuid" validate:"required"` // 必填
	rest string `json:"schedulerTriggerUuid" validate:"required"` // 必填
	rest bool `json:"triggerNow,omitempty"`
}

// AddSchedulerJobGroupToSchedulerTriggerParam AddSchedulerJobGroupToSchedulerTrigger请求参数
type AddSchedulerJobGroupToSchedulerTriggerParam struct {
	BaseParam
	Params AddSchedulerJobGroupToSchedulerTriggerDetailParam `json:"params"` // 详细参数
}

