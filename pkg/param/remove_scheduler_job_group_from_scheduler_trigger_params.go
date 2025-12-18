// Copyright (c) ZStack.io, Inc.

package param

// RemoveSchedulerJobGroupFromSchedulerTriggerDetailParam RemoveSchedulerJobGroupFromSchedulerTrigger detail param
type RemoveSchedulerJobGroupFromSchedulerTriggerDetailParam struct {
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid" validate:"required"`
	SchedulerTriggerUuid string `json:"schedulerTriggerUuid" validate:"required"`
}

// RemoveSchedulerJobGroupFromSchedulerTriggerParam RemoveSchedulerJobGroupFromSchedulerTrigger request param
type RemoveSchedulerJobGroupFromSchedulerTriggerParam struct {
	BaseParam
	Params RemoveSchedulerJobGroupFromSchedulerTriggerDetailParam `json:"params"`
}
