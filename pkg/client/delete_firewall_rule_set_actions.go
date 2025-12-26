// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteFirewallRuleSet deletes FirewallRuleSet
func (cli *ZSClient) DeleteFirewallRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/ruleSets/{uuid}", uuid, string(deleteMode))
}
