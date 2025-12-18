// Copyright (c) ZStack.io, Inc.

package param

// AttachPortForwardingRuleDetailParam AttachPortForwardingRule detail param
type AttachPortForwardingRuleDetailParam struct {
	RuleUuid string `json:"ruleUuid" validate:"required"`
	VmNicUuid string `json:"vmNicUuid" validate:"required"`
}

// AttachPortForwardingRuleParam AttachPortForwardingRule request param
type AttachPortForwardingRuleParam struct {
	BaseParam
	Params AttachPortForwardingRuleDetailParam `json:"params"`
}
