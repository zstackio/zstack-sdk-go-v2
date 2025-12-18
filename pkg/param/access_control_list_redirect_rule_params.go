// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccessControlListRedirectRuleDetailParam ChangeAccessControlListRedirectRule详细参数
type ChangeAccessControlListRedirectRuleDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
}

// ChangeAccessControlListRedirectRuleParam ChangeAccessControlListRedirectRule请求参数
type ChangeAccessControlListRedirectRuleParam struct {
	BaseParam
	Params ChangeAccessControlListRedirectRuleDetailParam `json:"params"` // 详细参数
}

