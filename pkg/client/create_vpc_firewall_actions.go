// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVpcFirewall creates VpcFirewall
func (cli *ZSClient) CreateVpcFirewall(params param.CreateVpcFirewallParam) (*view.CreateVpcFirewallEventView, error) {
	resp := view.CreateVpcFirewallEventView{}
	if err := cli.Post("v1/vpcfirewalls", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
