// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteFirewallRule deletes FirewallRule
func (cli *ZSClient) DeleteFirewallRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/rules/{uuid}", uuid, string(deleteMode))
}
