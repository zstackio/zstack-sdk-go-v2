// Copyright (c) ZStack.io, Inc.

package param

// RunSchedulerTriggerDetailParam RunSchedulerTrigger详细参数
type RunSchedulerTriggerDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"jobUuids,omitempty"`
}

// RunSchedulerTriggerParam RunSchedulerTrigger请求参数
type RunSchedulerTriggerParam struct {
	BaseParam
	Params RunSchedulerTriggerDetailParam `json:"params"` // 详细参数
}

