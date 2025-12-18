// Copyright (c) ZStack.io, Inc.

package param

// GetLdapEntryDetailParam GetLdapEntry detail param
type GetLdapEntryDetailParam struct {
	LdapFilter string `json:"ldapFilter" validate:"required"`
	Limit int `json:"limit,omitempty"`
	LdapServerUuid string `json:"ldapServerUuid,omitempty"`
}

// GetLdapEntryParam GetLdapEntry request param
type GetLdapEntryParam struct {
	BaseParam
	Params GetLdapEntryDetailParam `json:"params"`
}
