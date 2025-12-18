// Copyright (c) ZStack.io, Inc.

package param

// CleanInvalidLdapBindingDetailParam CleanInvalidLdapBinding detail param
type CleanInvalidLdapBindingDetailParam struct {
}

// CleanInvalidLdapBindingParam CleanInvalidLdapBinding request param
type CleanInvalidLdapBindingParam struct {
	BaseParam
	Params CleanInvalidLdapBindingDetailParam `json:"params"`
}
