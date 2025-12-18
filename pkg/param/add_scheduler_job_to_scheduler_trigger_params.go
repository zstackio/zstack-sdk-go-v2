// Copyright (c) ZStack.io, Inc.

package param

// AddSchedulerJobToSchedulerTriggerDetailParam AddSchedulerJobToSchedulerTrigger detail param
type AddSchedulerJobToSchedulerTriggerDetailParam struct {
	SchedulerJobUuid string `json:"schedulerJobUuid" validate:"required"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid" validate:"required"`
	TriggerNow bool `json:"triggerNow,omitempty"`
}

// AddSchedulerJobToSchedulerTriggerParam AddSchedulerJobToSchedulerTrigger request param
type AddSchedulerJobToSchedulerTriggerParam struct {
	BaseParam
	Params AddSchedulerJobToSchedulerTriggerDetailParam `json:"params"`
}
