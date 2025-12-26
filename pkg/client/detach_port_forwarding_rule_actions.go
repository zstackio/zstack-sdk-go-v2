// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) DetachPortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding/{uuid}/vm-instances/nics", uuid, string(deleteMode))
}
