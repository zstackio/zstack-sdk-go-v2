// Copyright (c) ZStack.io, Inc.

package param

// QueryZBoxDetailParam QueryZBox详细参数
type QueryZBoxDetailParam struct {
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

// QueryZBoxParam QueryZBox请求参数
type QueryZBoxParam struct {
	BaseParam
	Params QueryZBoxDetailParam `json:"params"` // 详细参数
}

