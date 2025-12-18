// Copyright (c) ZStack.io, Inc.

package param

// GetNoTriggerSchedulerJobsDetailParam GetNoTriggerSchedulerJobs详细参数
type GetNoTriggerSchedulerJobsDetailParam struct {
}

// GetNoTriggerSchedulerJobsParam GetNoTriggerSchedulerJobs请求参数
type GetNoTriggerSchedulerJobsParam struct {
	BaseParam
	Params GetNoTriggerSchedulerJobsDetailParam `json:"params"` // 详细参数
}

