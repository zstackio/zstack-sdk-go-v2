// Copyright (c) ZStack.io, Inc.

package param

// RemoveSchedulerJobsFromSchedulerJobGroupDetailParam RemoveSchedulerJobsFromSchedulerJobGroup详细参数
type RemoveSchedulerJobsFromSchedulerJobGroupDetailParam struct {
	rest string `json:"schedulerJobGroupUuid" validate:"required"` // 必填
	rest []string `json:"schedulerJobUuids" validate:"required"` // 必填
}

// RemoveSchedulerJobsFromSchedulerJobGroupParam RemoveSchedulerJobsFromSchedulerJobGroup请求参数
type RemoveSchedulerJobsFromSchedulerJobGroupParam struct {
	BaseParam
	Params RemoveSchedulerJobsFromSchedulerJobGroupDetailParam `json:"params"` // 详细参数
}

