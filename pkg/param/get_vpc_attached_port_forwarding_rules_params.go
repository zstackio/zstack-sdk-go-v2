// Copyright (c) ZStack.io, Inc.

package param

// GetVpcAttachedPortForwardingRulesDetailParam GetVpcAttachedPortForwardingRules detail param
type GetVpcAttachedPortForwardingRulesDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetVpcAttachedPortForwardingRulesParam GetVpcAttachedPortForwardingRules request param
type GetVpcAttachedPortForwardingRulesParam struct {
	BaseParam
	Params GetVpcAttachedPortForwardingRulesDetailParam `json:"params"`
}
