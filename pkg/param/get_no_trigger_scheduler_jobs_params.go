// Copyright (c) ZStack.io, Inc.

package param

// GetNoTriggerSchedulerJobsDetailParam GetNoTriggerSchedulerJobs detail param
type GetNoTriggerSchedulerJobsDetailParam struct {
}

// GetNoTriggerSchedulerJobsParam GetNoTriggerSchedulerJobs request param
type GetNoTriggerSchedulerJobsParam struct {
	BaseParam
	Params GetNoTriggerSchedulerJobsDetailParam `json:"params"`
}
