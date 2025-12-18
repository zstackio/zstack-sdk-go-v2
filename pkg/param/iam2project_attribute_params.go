// Copyright (c) ZStack.io, Inc.

package param

// QueryIAM2ProjectAttributeDetailParam QueryIAM2ProjectAttribute详细参数
type QueryIAM2ProjectAttributeDetailParam struct {
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

// QueryIAM2ProjectAttributeParam QueryIAM2ProjectAttribute请求参数
type QueryIAM2ProjectAttributeParam struct {
	BaseParam
	Params QueryIAM2ProjectAttributeDetailParam `json:"params"` // 详细参数
}

