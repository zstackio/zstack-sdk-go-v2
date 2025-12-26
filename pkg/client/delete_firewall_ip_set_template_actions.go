// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteFirewallIpSetTemplate deletes FirewallIpSetTemplate
func (cli *ZSClient) DeleteFirewallIpSetTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/ipset/templates/{uuid}", uuid, string(deleteMode))
}
