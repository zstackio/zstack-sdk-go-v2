// Copyright (c) ZStack.io, Inc.

package param

// AddLdapServerDetailParam AddLdapServer detail param
type AddLdapServerDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	Base string `json:"base" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Encryption string `json:"encryption" validate:"required"`
	Scope string `json:"scope" validate:"required"`
}

// AddLdapServerParam AddLdapServer request param
type AddLdapServerParam struct {
	BaseParam
	Params AddLdapServerDetailParam `json:"params"`
}
