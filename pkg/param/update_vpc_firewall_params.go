// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcFirewallDetailParam UpdateVpcFirewall detail param
type UpdateVpcFirewallDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateVpcFirewallParam UpdateVpcFirewall request param
type UpdateVpcFirewallParam struct {
	BaseParam
	Params UpdateVpcFirewallDetailParam `json:"params"`
}
