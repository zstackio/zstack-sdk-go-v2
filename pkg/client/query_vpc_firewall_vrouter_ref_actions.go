// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcFirewallVRouterRef queries VpcFirewallVRouterRef list
func (cli *ZSClient) QueryVpcFirewallVRouterRef(params param.QueryParam) ([]view.VpcFirewallVRouterRefInventoryView, error) {
	var resp []view.VpcFirewallVRouterRefInventoryView
	return resp, cli.List("v1/vpcfirewalls/vrouters/refs", &params, &resp)
}
