// Copyright (c) ZStack.io, Inc.

package param

// QueryIAM2LdapBindingDetailParam QueryIAM2LdapBinding详细参数
type QueryIAM2LdapBindingDetailParam struct {
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

// QueryIAM2LdapBindingParam QueryIAM2LdapBinding请求参数
type QueryIAM2LdapBindingParam struct {
	BaseParam
	Params QueryIAM2LdapBindingDetailParam `json:"params"` // 详细参数
}

