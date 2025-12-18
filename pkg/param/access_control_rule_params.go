// Copyright (c) ZStack.io, Inc.

package param

// UpdateAccessControlRuleDetailParam UpdateAccessControlRule详细参数
type UpdateAccessControlRuleDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"rule,omitempty"`
}

// UpdateAccessControlRuleParam UpdateAccessControlRule请求参数
type UpdateAccessControlRuleParam struct {
	BaseParam
	Params UpdateAccessControlRuleDetailParam `json:"params"` // 详细参数
}

