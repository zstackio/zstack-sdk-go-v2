// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) DetachPortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding/{uuid}/vm-instances/nics", uuid, string(deleteMode))
}
