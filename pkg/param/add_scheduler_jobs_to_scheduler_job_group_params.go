// Copyright (c) ZStack.io, Inc.

package param

// AddSchedulerJobsToSchedulerJobGroupDetailParam AddSchedulerJobsToSchedulerJobGroup详细参数
type AddSchedulerJobsToSchedulerJobGroupDetailParam struct {
	rest string `json:"schedulerJobGroupUuid" validate:"required"` // 必填
	rest []string `json:"schedulerJobUuids" validate:"required"` // 必填
}

// AddSchedulerJobsToSchedulerJobGroupParam AddSchedulerJobsToSchedulerJobGroup请求参数
type AddSchedulerJobsToSchedulerJobGroupParam struct {
	BaseParam
	Params AddSchedulerJobsToSchedulerJobGroupDetailParam `json:"params"` // 详细参数
}

