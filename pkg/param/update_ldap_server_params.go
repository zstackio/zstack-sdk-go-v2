// Copyright (c) ZStack.io, Inc.

package param

// UpdateLdapServerDetailParam UpdateLdapServer detail param
type UpdateLdapServerDetailParam struct {
	LdapServerUuid string `json:"ldapServerUuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	Base string `json:"base,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Encryption string `json:"encryption,omitempty"`
}

// UpdateLdapServerParam UpdateLdapServer request param
type UpdateLdapServerParam struct {
	BaseParam
	Params UpdateLdapServerDetailParam `json:"params"`
}
