// Copyright (c) ZStack.io, Inc.

package param

// GetZWatchAlertHistogramDetailParam GetZWatchAlertHistogram详细参数
type GetZWatchAlertHistogramDetailParam struct {
	rest string `json:"tableName" validate:"required"` // 必填
	rest int64 `json:"startTime" validate:"required"` // 必填
	rest int64 `json:"endTime" validate:"required"` // 必填
	rest int `json:"intervalHours" validate:"required"` // 必填
	rest []string `json:"groupColumns,omitempty"`
}

// GetZWatchAlertHistogramParam GetZWatchAlertHistogram请求参数
type GetZWatchAlertHistogramParam struct {
	BaseParam
	Params GetZWatchAlertHistogramDetailParam `json:"params"` // 详细参数
}

