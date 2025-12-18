// Copyright (c) ZStack.io, Inc.

package param

// GetSchedulerExecutionReportDetailParam GetSchedulerExecutionReport detail param
type GetSchedulerExecutionReportDetailParam struct {
	StartTime int64 `json:"startTime" validate:"required"`
	IntervalTimeUnit string `json:"intervalTimeUnit" validate:"required"`
	Range int `json:"range" validate:"required"`
	SchedulerJobTypes []string `json:"schedulerJobTypes" validate:"required"`
}

// GetSchedulerExecutionReportParam GetSchedulerExecutionReport request param
type GetSchedulerExecutionReportParam struct {
	BaseParam
	Params GetSchedulerExecutionReportDetailParam `json:"params"`
}
