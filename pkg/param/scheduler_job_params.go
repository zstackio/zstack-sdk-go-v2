// Copyright (c) ZStack.io, Inc.

package param

// DeleteSchedulerJobDetailParam DeleteSchedulerJob详细参数
type DeleteSchedulerJobDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteSchedulerJobParam DeleteSchedulerJob请求参数
type DeleteSchedulerJobParam struct {
	BaseParam
	Params DeleteSchedulerJobDetailParam `json:"params"` // 详细参数
}

