// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallDetailParam DeleteFirewall详细参数
type DeleteFirewallDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteFirewallParam DeleteFirewall请求参数
type DeleteFirewallParam struct {
	BaseParam
	Params DeleteFirewallDetailParam `json:"params"` // 详细参数
}

