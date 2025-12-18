// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateLdapEntryForBindingDetailParam GetCandidateLdapEntryForBinding detail param
type GetCandidateLdapEntryForBindingDetailParam struct {
	LdapFilter string `json:"ldapFilter" validate:"required"`
	Limit int `json:"limit,omitempty"`
}

// GetCandidateLdapEntryForBindingParam GetCandidateLdapEntryForBinding request param
type GetCandidateLdapEntryForBindingParam struct {
	BaseParam
	Params GetCandidateLdapEntryForBindingDetailParam `json:"params"`
}
