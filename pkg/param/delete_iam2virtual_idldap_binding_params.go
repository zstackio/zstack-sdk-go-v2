// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2VirtualIDLdapBindingDetailParam DeleteIAM2VirtualIDLdapBinding detail param
type DeleteIAM2VirtualIDLdapBindingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIAM2VirtualIDLdapBindingParam DeleteIAM2VirtualIDLdapBinding request param
type DeleteIAM2VirtualIDLdapBindingParam struct {
	BaseParam
	Params DeleteIAM2VirtualIDLdapBindingDetailParam `json:"params"`
}
