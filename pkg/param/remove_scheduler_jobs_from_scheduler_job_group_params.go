// Copyright (c) ZStack.io, Inc.

package param

// RemoveSchedulerJobsFromSchedulerJobGroupDetailParam RemoveSchedulerJobsFromSchedulerJobGroup detail param
type RemoveSchedulerJobsFromSchedulerJobGroupDetailParam struct {
	SchedulerJobGroupUuid string `json:"schedulerJobGroupUuid" validate:"required"`
	SchedulerJobUuids []string `json:"schedulerJobUuids" validate:"required"`
}

// RemoveSchedulerJobsFromSchedulerJobGroupParam RemoveSchedulerJobsFromSchedulerJobGroup request param
type RemoveSchedulerJobsFromSchedulerJobGroupParam struct {
	BaseParam
	Params RemoveSchedulerJobsFromSchedulerJobGroupDetailParam `json:"params"`
}
