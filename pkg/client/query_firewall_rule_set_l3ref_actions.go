// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFirewallRuleSetL3Ref queries FirewallRuleSetL3Ref list
func (cli *ZSClient) QueryFirewallRuleSetL3Ref(params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp []view.VpcFirewallRuleSetL3RefInventoryView
	return resp, cli.List("v1/vpcfirewalls/l3networks/rulesets/refs", params, &resp)
}
