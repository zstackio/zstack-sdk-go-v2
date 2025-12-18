// Copyright (c) ZStack.io, Inc.

package param

// CreateVpcFirewallDetailParam CreateVpcFirewall detail param
type CreateVpcFirewallDetailParam struct {
	VpcUuid string `json:"vpcUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVpcFirewallParam CreateVpcFirewall request param
type CreateVpcFirewallParam struct {
	BaseParam
	Params CreateVpcFirewallDetailParam `json:"params"`
}
