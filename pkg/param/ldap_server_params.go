// Copyright (c) ZStack.io, Inc.

package param

// QueryLdapServerDetailParam QueryLdapServer详细参数
type QueryLdapServerDetailParam struct {
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

// QueryLdapServerParam QueryLdapServer请求参数
type QueryLdapServerParam struct {
	BaseParam
	Params QueryLdapServerDetailParam `json:"params"` // 详细参数
}

