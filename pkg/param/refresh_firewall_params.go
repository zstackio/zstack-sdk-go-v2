// Copyright (c) ZStack.io, Inc.

package param

// RefreshFirewallDetailParam RefreshFirewall detail param
type RefreshFirewallDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RefreshFirewallParam RefreshFirewall request param
type RefreshFirewallParam struct {
	BaseParam
	Params RefreshFirewallDetailParam `json:"params"`
}
