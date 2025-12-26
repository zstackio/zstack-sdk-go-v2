// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteFirewallRule deletes FirewallRule
func (cli *ZSClient) DeleteFirewallRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/rules/{uuid}", uuid, string(deleteMode))
}
