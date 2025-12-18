// Copyright (c) ZStack.io, Inc.

package param

// UpdateSchedulerJobGroupDetailParam UpdateSchedulerJobGroup详细参数
type UpdateSchedulerJobGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest map[string]string `json:"parameters,omitempty"`
}

// UpdateSchedulerJobGroupParam UpdateSchedulerJobGroup请求参数
type UpdateSchedulerJobGroupParam struct {
	BaseParam
	Params UpdateSchedulerJobGroupDetailParam `json:"params"` // 详细参数
}

