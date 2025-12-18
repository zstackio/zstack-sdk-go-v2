// Copyright (c) ZStack.io, Inc.

package param

// ChangeAccessControlListRedirectRuleDetailParam ChangeAccessControlListRedirectRule detail param
type ChangeAccessControlListRedirectRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
}

// ChangeAccessControlListRedirectRuleParam ChangeAccessControlListRedirectRule request param
type ChangeAccessControlListRedirectRuleParam struct {
	BaseParam
	Params ChangeAccessControlListRedirectRuleDetailParam `json:"params"`
}
