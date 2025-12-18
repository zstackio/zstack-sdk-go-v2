// Copyright (c) ZStack.io, Inc.

package param

// DeletePortForwardingRuleDetailParam DeletePortForwardingRule detail param
type DeletePortForwardingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePortForwardingRuleParam DeletePortForwardingRule request param
type DeletePortForwardingRuleParam struct {
	BaseParam
	Params DeletePortForwardingRuleDetailParam `json:"params"`
}
