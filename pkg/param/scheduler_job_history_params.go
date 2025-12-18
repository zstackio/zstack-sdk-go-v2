// Copyright (c) ZStack.io, Inc.

package param

// QuerySchedulerJobHistoryDetailParam QuerySchedulerJobHistory详细参数
type QuerySchedulerJobHistoryDetailParam struct {
	rest []interface{} `json:"conditions" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
	rest bool `json:"count,omitempty"`
	rest string `json:"groupBy,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
	rest string `json:"filterName,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest []string `json:"fields,omitempty"`
}

// QuerySchedulerJobHistoryParam QuerySchedulerJobHistory请求参数
type QuerySchedulerJobHistoryParam struct {
	BaseParam
	Params QuerySchedulerJobHistoryDetailParam `json:"params"` // 详细参数
}

