// Copyright (c) ZStack.io, Inc.

package param

// CreateSchedulerTriggerDetailParam CreateSchedulerTrigger详细参数
type CreateSchedulerTriggerDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest int `json:"schedulerInterval,omitempty"`
	rest int `json:"repeatCount,omitempty"`
	rest int64 `json:"startTime,omitempty"`
	rest string `json:"schedulerType" validate:"required"` // 必填
	rest string `json:"cron,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSchedulerTriggerParam CreateSchedulerTrigger请求参数
type CreateSchedulerTriggerParam struct {
	BaseParam
	Params CreateSchedulerTriggerDetailParam `json:"params"` // 详细参数
}

