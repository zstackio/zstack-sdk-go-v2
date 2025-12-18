// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2VirtualIDLdapBindingDetailParam CreateIAM2VirtualIDLdapBinding detail param
type CreateIAM2VirtualIDLdapBindingDetailParam struct {
	VirtualIDUuid string `json:"virtualIDUuid" validate:"required"`
	LdapUid string `json:"ldapUid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDLdapBindingParam CreateIAM2VirtualIDLdapBinding request param
type CreateIAM2VirtualIDLdapBindingParam struct {
	BaseParam
	Params CreateIAM2VirtualIDLdapBindingDetailParam `json:"params"`
}
