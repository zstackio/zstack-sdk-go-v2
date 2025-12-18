// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcFirewall queries VpcFirewall list
func (cli *ZSClient) QueryVpcFirewall(params param.QueryParam) ([]view.VpcFirewallInventoryView, error) {
	var resp []view.VpcFirewallInventoryView
	return resp, cli.List("v1/vpcfirewalls", &params, &resp)
}
