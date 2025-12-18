// Copyright (c) ZStack.io, Inc.

package param

// UpdatePortForwardingRuleDetailParam UpdatePortForwardingRule详细参数
type UpdatePortForwardingRuleDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdatePortForwardingRuleParam UpdatePortForwardingRule请求参数
type UpdatePortForwardingRuleParam struct {
	BaseParam
	Params UpdatePortForwardingRuleDetailParam `json:"params"` // 详细参数
}

