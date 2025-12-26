// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFirewallRuleTemplate queries FirewallRuleTemplate list
func (cli *ZSClient) QueryFirewallRuleTemplate(params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp []view.VpcFirewallRuleTemplateInventoryView
	return resp, cli.List("v1/vpcfirewalls/rules/templates", params, &resp)
}
