// Copyright (c) ZStack.io, Inc.

package param

// ChangePortForwardingRuleStateDetailParam ChangePortForwardingRuleState detail param
type ChangePortForwardingRuleStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangePortForwardingRuleStateParam ChangePortForwardingRuleState request param
type ChangePortForwardingRuleStateParam struct {
	BaseParam
	Params ChangePortForwardingRuleStateDetailParam `json:"params"`
}
