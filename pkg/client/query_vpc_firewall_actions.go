// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcFirewall queries VpcFirewall list
func (cli *ZSClient) QueryVpcFirewall(params *param.QueryParam) ([]view.VpcFirewallInventoryView, error) {
	var resp []view.VpcFirewallInventoryView
	return resp, cli.List("v1/vpcfirewalls", params, &resp)
}
