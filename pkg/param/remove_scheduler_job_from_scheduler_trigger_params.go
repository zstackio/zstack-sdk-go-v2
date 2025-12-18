// Copyright (c) ZStack.io, Inc.

package param

// RemoveSchedulerJobFromSchedulerTriggerDetailParam RemoveSchedulerJobFromSchedulerTrigger detail param
type RemoveSchedulerJobFromSchedulerTriggerDetailParam struct {
	SchedulerJobUuid string `json:"schedulerJobUuid" validate:"required"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid" validate:"required"`
}

// RemoveSchedulerJobFromSchedulerTriggerParam RemoveSchedulerJobFromSchedulerTrigger request param
type RemoveSchedulerJobFromSchedulerTriggerParam struct {
	BaseParam
	Params RemoveSchedulerJobFromSchedulerTriggerDetailParam `json:"params"`
}
