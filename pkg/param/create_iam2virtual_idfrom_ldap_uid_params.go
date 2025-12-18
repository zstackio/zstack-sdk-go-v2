// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2VirtualIDFromLdapUidDetailParam CreateIAM2VirtualIDFromLdapUid detail param
type CreateIAM2VirtualIDFromLdapUidDetailParam struct {
	LdapUid string `json:"ldapUid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2VirtualIDFromLdapUidParam CreateIAM2VirtualIDFromLdapUid request param
type CreateIAM2VirtualIDFromLdapUidParam struct {
	BaseParam
	Params CreateIAM2VirtualIDFromLdapUidDetailParam `json:"params"`
}
