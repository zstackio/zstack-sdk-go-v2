// Copyright (c) ZStack.io, Inc.

package param

// DeleteLdapServerDetailParam DeleteLdapServer detail param
type DeleteLdapServerDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteLdapServerParam DeleteLdapServer request param
type DeleteLdapServerParam struct {
	BaseParam
	Params DeleteLdapServerDetailParam `json:"params"`
}
