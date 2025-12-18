// Copyright (c) ZStack.io, Inc.

package param

// DetachPortForwardingRuleDetailParam DetachPortForwardingRule detail param
type DetachPortForwardingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DetachPortForwardingRuleParam DetachPortForwardingRule request param
type DetachPortForwardingRuleParam struct {
	BaseParam
	Params DetachPortForwardingRuleDetailParam `json:"params"`
}
