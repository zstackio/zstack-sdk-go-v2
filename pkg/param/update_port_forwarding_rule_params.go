// Copyright (c) ZStack.io, Inc.

package param

// UpdatePortForwardingRuleDetailParam UpdatePortForwardingRule detail param
type UpdatePortForwardingRuleDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdatePortForwardingRuleParam UpdatePortForwardingRule request param
type UpdatePortForwardingRuleParam struct {
	BaseParam
	Params UpdatePortForwardingRuleDetailParam `json:"params"`
}
