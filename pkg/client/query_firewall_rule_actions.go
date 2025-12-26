// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFirewallRule queries FirewallRule list
func (cli *ZSClient) QueryFirewallRule(params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, error) {
	var resp []view.VpcFirewallRuleInventoryView
	return resp, cli.List("v1/vpcfirewalls/rules", params, &resp)
}
