// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcFirewallVRouterRef queries VpcFirewallVRouterRef list
func (cli *ZSClient) QueryVpcFirewallVRouterRef(params *param.QueryParam) ([]view.VpcFirewallVRouterRefInventoryView, error) {
	var resp []view.VpcFirewallVRouterRefInventoryView
	return resp, cli.List("v1/vpcfirewalls/vrouters/refs", params, &resp)
}
