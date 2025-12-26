// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFirewallIpSetTemplate queries FirewallIpSetTemplate list
func (cli *ZSClient) QueryFirewallIpSetTemplate(params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp []view.VpcFirewallIpSetTemplateInventoryView
	return resp, cli.List("v1/vpcfirewalls/ipset/templates", params, &resp)
}
