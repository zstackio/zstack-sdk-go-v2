// Copyright (c) ZStack.io, Inc.

package param

// AddAccessControlListRedirectRuleDetailParam AddAccessControlListRedirectRule详细参数
type AddAccessControlListRedirectRuleDetailParam struct {
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"domain,omitempty"`
	rest string `json:"url,omitempty"`
	rest string `json:"aclUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAccessControlListRedirectRuleParam AddAccessControlListRedirectRule请求参数
type AddAccessControlListRedirectRuleParam struct {
	BaseParam
	Params AddAccessControlListRedirectRuleDetailParam `json:"params"` // 详细参数
}

