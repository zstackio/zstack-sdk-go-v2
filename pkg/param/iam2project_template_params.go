// Copyright (c) ZStack.io, Inc.

package param

// QueryIAM2ProjectTemplateDetailParam QueryIAM2ProjectTemplate详细参数
type QueryIAM2ProjectTemplateDetailParam struct {
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

// QueryIAM2ProjectTemplateParam QueryIAM2ProjectTemplate请求参数
type QueryIAM2ProjectTemplateParam struct {
	BaseParam
	Params QueryIAM2ProjectTemplateDetailParam `json:"params"` // 详细参数
}

