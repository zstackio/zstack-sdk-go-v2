// Copyright (c) ZStack.io, Inc.

package param

// RemoveSchedulerJobGroupFromSchedulerTriggerDetailParam RemoveSchedulerJobGroupFromSchedulerTrigger详细参数
type RemoveSchedulerJobGroupFromSchedulerTriggerDetailParam struct {
	rest string `json:"schedulerJobGroupUuid" validate:"required"` // 必填
	rest string `json:"schedulerTriggerUuid" validate:"required"` // 必填
}

// RemoveSchedulerJobGroupFromSchedulerTriggerParam RemoveSchedulerJobGroupFromSchedulerTrigger请求参数
type RemoveSchedulerJobGroupFromSchedulerTriggerParam struct {
	BaseParam
	Params RemoveSchedulerJobGroupFromSchedulerTriggerDetailParam `json:"params"` // 详细参数
}

