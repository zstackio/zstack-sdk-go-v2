// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateLdapEntryForIAM2BindingDetailParam GetCandidateLdapEntryForIAM2Binding详细参数
type GetCandidateLdapEntryForIAM2BindingDetailParam struct {
	rest string `json:"ldapFilter" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
}

// GetCandidateLdapEntryForIAM2BindingParam GetCandidateLdapEntryForIAM2Binding请求参数
type GetCandidateLdapEntryForIAM2BindingParam struct {
	BaseParam
	Params GetCandidateLdapEntryForIAM2BindingDetailParam `json:"params"` // 详细参数
}

