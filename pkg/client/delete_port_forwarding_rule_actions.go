// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePortForwardingRule deletes PortForwardingRule
func (cli *ZSClient) DeletePortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding/{uuid}", uuid, string(deleteMode))
}
