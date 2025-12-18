// Copyright (c) ZStack.io, Inc.

package param

// CreateLdapBindingDetailParam CreateLdapBinding detail param
type CreateLdapBindingDetailParam struct {
	LdapUid string `json:"ldapUid" validate:"required"`
	AccountUuid string `json:"accountUuid" validate:"required"`
}

// CreateLdapBindingParam CreateLdapBinding request param
type CreateLdapBindingParam struct {
	BaseParam
	Params CreateLdapBindingDetailParam `json:"params"`
}
