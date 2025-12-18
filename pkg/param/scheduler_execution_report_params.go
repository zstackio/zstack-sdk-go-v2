// Copyright (c) ZStack.io, Inc.

package param

// GetSchedulerExecutionReportDetailParam GetSchedulerExecutionReport详细参数
type GetSchedulerExecutionReportDetailParam struct {
	rest int64 `json:"startTime" validate:"required"` // 必填
	rest string `json:"intervalTimeUnit" validate:"required"` // 必填
	rest int `json:"range" validate:"required"` // 必填
	rest []string `json:"schedulerJobTypes" validate:"required"` // 必填
}

// GetSchedulerExecutionReportParam GetSchedulerExecutionReport请求参数
type GetSchedulerExecutionReportParam struct {
	BaseParam
	Params GetSchedulerExecutionReportDetailParam `json:"params"` // 详细参数
}

