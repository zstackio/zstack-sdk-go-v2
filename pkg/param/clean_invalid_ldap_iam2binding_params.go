// Copyright (c) ZStack.io, Inc.

package param

// CleanInvalidLdapIAM2BindingDetailParam CleanInvalidLdapIAM2Binding detail param
type CleanInvalidLdapIAM2BindingDetailParam struct {
}

// CleanInvalidLdapIAM2BindingParam CleanInvalidLdapIAM2Binding request param
type CleanInvalidLdapIAM2BindingParam struct {
	BaseParam
	Params CleanInvalidLdapIAM2BindingDetailParam `json:"params"`
}
