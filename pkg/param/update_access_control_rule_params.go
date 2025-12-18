// Copyright (c) ZStack.io, Inc.

package param

// UpdateAccessControlRuleDetailParam UpdateAccessControlRule detail param
type UpdateAccessControlRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Rule string `json:"rule,omitempty"`
}

// UpdateAccessControlRuleParam UpdateAccessControlRule request param
type UpdateAccessControlRuleParam struct {
	BaseParam
	Params UpdateAccessControlRuleDetailParam `json:"params"`
}
