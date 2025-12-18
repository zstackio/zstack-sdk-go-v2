// Copyright (c) ZStack.io, Inc.

package param

// DeleteLdapBindingDetailParam DeleteLdapBinding detail param
type DeleteLdapBindingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteLdapBindingParam DeleteLdapBinding request param
type DeleteLdapBindingParam struct {
	BaseParam
	Params DeleteLdapBindingDetailParam `json:"params"`
}
