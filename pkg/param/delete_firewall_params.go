// Copyright (c) ZStack.io, Inc.

package param

// DeleteFirewallDetailParam DeleteFirewall detail param
type DeleteFirewallDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteFirewallParam DeleteFirewall request param
type DeleteFirewallParam struct {
	BaseParam
	Params DeleteFirewallDetailParam `json:"params"`
}
