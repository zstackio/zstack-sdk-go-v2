// Copyright (c) ZStack.io, Inc.

package param

// ChangePortForwardingRuleStateDetailParam ChangePortForwardingRuleState详细参数
type ChangePortForwardingRuleStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangePortForwardingRuleStateParam ChangePortForwardingRuleState请求参数
type ChangePortForwardingRuleStateParam struct {
	BaseParam
	Params ChangePortForwardingRuleStateDetailParam `json:"params"` // 详细参数
}

