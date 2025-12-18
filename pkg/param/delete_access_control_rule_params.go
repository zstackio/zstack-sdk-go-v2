// Copyright (c) ZStack.io, Inc.

package param

// DeleteAccessControlRuleDetailParam DeleteAccessControlRule detail param
type DeleteAccessControlRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAccessControlRuleParam DeleteAccessControlRule request param
type DeleteAccessControlRuleParam struct {
	BaseParam
	Params DeleteAccessControlRuleDetailParam `json:"params"`
}
