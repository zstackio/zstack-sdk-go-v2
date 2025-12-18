// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpcFirewall creates VpcFirewall
func (cli *ZSClient) CreateVpcFirewall(params param.CreateVpcFirewallParam) (*view.CreateVpcFirewallEventView, error) {
	resp := view.CreateVpcFirewallEventView{}
	if err := cli.Post("v1/vpcfirewalls", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
