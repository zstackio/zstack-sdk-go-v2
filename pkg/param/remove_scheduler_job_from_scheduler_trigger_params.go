// Copyright (c) ZStack.io, Inc.

package param

// RemoveSchedulerJobFromSchedulerTriggerDetailParam RemoveSchedulerJobFromSchedulerTrigger详细参数
type RemoveSchedulerJobFromSchedulerTriggerDetailParam struct {
	rest string `json:"schedulerJobUuid" validate:"required"` // 必填
	rest string `json:"schedulerTriggerUuid" validate:"required"` // 必填
}

// RemoveSchedulerJobFromSchedulerTriggerParam RemoveSchedulerJobFromSchedulerTrigger请求参数
type RemoveSchedulerJobFromSchedulerTriggerParam struct {
	BaseParam
	Params RemoveSchedulerJobFromSchedulerTriggerDetailParam `json:"params"` // 详细参数
}

