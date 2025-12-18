// Copyright (c) ZStack.io, Inc.

package param

// RefreshFirewallDetailParam RefreshFirewall详细参数
type RefreshFirewallDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RefreshFirewallParam RefreshFirewall请求参数
type RefreshFirewallParam struct {
	BaseParam
	Params RefreshFirewallDetailParam `json:"params"` // 详细参数
}

