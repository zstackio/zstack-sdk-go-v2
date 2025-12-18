// Copyright (c) ZStack.io, Inc.

package param

// QueryIAM2OrganizationDetailParam QueryIAM2Organization详细参数
type QueryIAM2OrganizationDetailParam struct {
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

// QueryIAM2OrganizationParam QueryIAM2Organization请求参数
type QueryIAM2OrganizationParam struct {
	BaseParam
	Params QueryIAM2OrganizationDetailParam `json:"params"` // 详细参数
}

