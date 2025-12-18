// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedPortForwardingRulesDetailParam GetVpcAttachedPortForwardingRules详细参数
type GetVpcAttachedPortForwardingRulesDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetVpcAttachedPortForwardingRulesParam GetVpcAttachedPortForwardingRules请求参数
type GetVpcAttachedPortForwardingRulesParam struct {
	BaseParam
	Params GetVpcAttachedPortForwardingRulesDetailParam `json:"params"` // 详细参数
}

