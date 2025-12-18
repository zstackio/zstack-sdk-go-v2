// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2VirtualIDFromLdapUidDetailParam CreateIAM2VirtualIDFromLdapUid详细参数
type CreateIAM2VirtualIDFromLdapUidDetailParam struct {
	rest string `json:"ldapUid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDFromLdapUidParam CreateIAM2VirtualIDFromLdapUid请求参数
type CreateIAM2VirtualIDFromLdapUidParam struct {
	BaseParam
	Params CreateIAM2VirtualIDFromLdapUidDetailParam `json:"params"` // 详细参数
}

