// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteFirewallRuleTemplate deletes FirewallRuleTemplate
func (cli *ZSClient) DeleteFirewallRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/rules/templates/{uuid}", uuid, string(deleteMode))
}
