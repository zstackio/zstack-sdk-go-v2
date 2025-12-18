// Copyright (c) ZStack.io, Inc.

package param

// AddSchedulerJobsToSchedulerJobGroupDetailParam AddSchedulerJobsToSchedulerJobGroup detail param
type AddSchedulerJobsToSchedulerJobGroupDetailParam struct {
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid" validate:"required"`
	SchedulerJobUuids []string `json:"schedulerJobUuids" validate:"required"`
}

// AddSchedulerJobsToSchedulerJobGroupParam AddSchedulerJobsToSchedulerJobGroup request param
type AddSchedulerJobsToSchedulerJobGroupParam struct {
	BaseParam
	Params AddSchedulerJobsToSchedulerJobGroupDetailParam `json:"params"`
}
