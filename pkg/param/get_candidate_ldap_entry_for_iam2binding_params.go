// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateLdapEntryForIAM2BindingDetailParam GetCandidateLdapEntryForIAM2Binding detail param
type GetCandidateLdapEntryForIAM2BindingDetailParam struct {
	LdapFilter string `json:"ldapFilter" validate:"required"`
	Limit int `json:"limit,omitempty"`
}

// GetCandidateLdapEntryForIAM2BindingParam GetCandidateLdapEntryForIAM2Binding request param
type GetCandidateLdapEntryForIAM2BindingParam struct {
	BaseParam
	Params GetCandidateLdapEntryForIAM2BindingDetailParam `json:"params"`
}
