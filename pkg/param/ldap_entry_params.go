// Copyright (c) ZStack.io, Inc.

package param

// GetLdapEntryDetailParam GetLdapEntry详细参数
type GetLdapEntryDetailParam struct {
	rest string `json:"ldapFilter" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest string `json:"ldapServerUuid,omitempty"`
}

// GetLdapEntryParam GetLdapEntry请求参数
type GetLdapEntryParam struct {
	BaseParam
	Params GetLdapEntryDetailParam `json:"params"` // 详细参数
}

