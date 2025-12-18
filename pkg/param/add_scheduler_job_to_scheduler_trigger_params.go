// Copyright (c) ZStack.io, Inc.

package param

// AddSchedulerJobToSchedulerTriggerDetailParam AddSchedulerJobToSchedulerTrigger详细参数
type AddSchedulerJobToSchedulerTriggerDetailParam struct {
	rest string `json:"schedulerJobUuid" validate:"required"` // 必填
	rest string `json:"schedulerTriggerUuid" validate:"required"` // 必填
	rest bool `json:"triggerNow,omitempty"`
}

// AddSchedulerJobToSchedulerTriggerParam AddSchedulerJobToSchedulerTrigger请求参数
type AddSchedulerJobToSchedulerTriggerParam struct {
	BaseParam
	Params AddSchedulerJobToSchedulerTriggerDetailParam `json:"params"` // 详细参数
}

