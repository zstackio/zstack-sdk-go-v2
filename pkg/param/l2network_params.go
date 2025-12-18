// Copyright (c) ZStack.io, Inc.

package param

// QueryL2NetworkDetailParam QueryL2Network详细参数
type QueryL2NetworkDetailParam struct {
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

// QueryL2NetworkParam QueryL2Network请求参数
type QueryL2NetworkParam struct {
	BaseParam
	Params QueryL2NetworkDetailParam `json:"params"` // 详细参数
}

