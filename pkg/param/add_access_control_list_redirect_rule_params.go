// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlListRedirectRuleDetailParam AddAccessControlListRedirectRule detail param
type AddAccessControlListRedirectRuleDetailParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Domain string `json:"domain,omitempty"`
	Url string `json:"url,omitempty"`
	AclUuid string `json:"aclUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAccessControlListRedirectRuleParam AddAccessControlListRedirectRule request param
type AddAccessControlListRedirectRuleParam struct {
	BaseParam
	Params AddAccessControlListRedirectRuleDetailParam `json:"params"`
}
