// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateLdapEntryForBindingDetailParam GetCandidateLdapEntryForBinding详细参数
type GetCandidateLdapEntryForBindingDetailParam struct {
	rest string `json:"ldapFilter" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
}

// GetCandidateLdapEntryForBindingParam GetCandidateLdapEntryForBinding请求参数
type GetCandidateLdapEntryForBindingParam struct {
	BaseParam
	Params GetCandidateLdapEntryForBindingDetailParam `json:"params"` // 详细参数
}

