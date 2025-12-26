// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFirewallRuleSet queries FirewallRuleSet list
func (cli *ZSClient) QueryFirewallRuleSet(params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, error) {
	var resp []view.VpcFirewallRuleSetInventoryView
	return resp, cli.List("v1/vpcfirewalls/ruleSets", params, &resp)
}
