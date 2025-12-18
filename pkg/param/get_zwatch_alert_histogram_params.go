// Copyright (c) ZStack.io, Inc.

package param

// GetZWatchAlertHistogramDetailParam GetZWatchAlertHistogram detail param
type GetZWatchAlertHistogramDetailParam struct {
	TableName string `json:"tableName" validate:"required"`
	StartTime int64 `json:"startTime" validate:"required"`
	EndTime int64 `json:"endTime" validate:"required"`
	IntervalHours int `json:"intervalHours" validate:"required"`
	GroupColumns []string `json:"groupColumns,omitempty"`
}

// GetZWatchAlertHistogramParam GetZWatchAlertHistogram request param
type GetZWatchAlertHistogramParam struct {
	BaseParam
	Params GetZWatchAlertHistogramDetailParam `json:"params"`
}
