// Copyright (c) ZStack.io, Inc.

package param

// QueryDRSAdviceDetailParam QueryDRSAdvice详细参数
type QueryDRSAdviceDetailParam struct {
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

// QueryDRSAdviceParam QueryDRSAdvice请求参数
type QueryDRSAdviceParam struct {
	BaseParam
	Params QueryDRSAdviceDetailParam `json:"params"` // 详细参数
}

