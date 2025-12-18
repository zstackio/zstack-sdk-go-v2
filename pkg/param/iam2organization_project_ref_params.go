// Copyright (c) ZStack.io, Inc.

package param

// QueryIAM2OrganizationProjectRefDetailParam QueryIAM2OrganizationProjectRef详细参数
type QueryIAM2OrganizationProjectRefDetailParam struct {
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

// QueryIAM2OrganizationProjectRefParam QueryIAM2OrganizationProjectRef请求参数
type QueryIAM2OrganizationProjectRefParam struct {
	BaseParam
	Params QueryIAM2OrganizationProjectRefDetailParam `json:"params"` // 详细参数
}

