// Copyright (c) ZStack.io, Inc.

package param

// AddSchedulerJobGroupToSchedulerTriggerDetailParam AddSchedulerJobGroupToSchedulerTrigger detail param
type AddSchedulerJobGroupToSchedulerTriggerDetailParam struct {
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid" validate:"required"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid" validate:"required"`
	TriggerNow bool `json:"triggerNow,omitempty"`
}

// AddSchedulerJobGroupToSchedulerTriggerParam AddSchedulerJobGroupToSchedulerTrigger request param
type AddSchedulerJobGroupToSchedulerTriggerParam struct {
	BaseParam
	Params AddSchedulerJobGroupToSchedulerTriggerDetailParam `json:"params"`
}
