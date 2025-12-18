// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2VirtualIDLdapBindingDetailParam CreateIAM2VirtualIDLdapBinding详细参数
type CreateIAM2VirtualIDLdapBindingDetailParam struct {
	rest string `json:"virtualIDUuid" validate:"required"` // 必填
	rest string `json:"ldapUid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDLdapBindingParam CreateIAM2VirtualIDLdapBinding请求参数
type CreateIAM2VirtualIDLdapBindingParam struct {
	BaseParam
	Params CreateIAM2VirtualIDLdapBindingDetailParam `json:"params"` // 详细参数
}

