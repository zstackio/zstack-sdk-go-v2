// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVpcFirewall updates VpcFirewall
func (cli *ZSClient) UpdateVpcFirewall(uuid string, params param.UpdateVpcFirewallParam) (*view.UpdateVpcFirewallEventView, error) {
	resp := view.UpdateVpcFirewallEventView{}
	if err := cli.Put("v1/vpcfirewalls/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
