// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePortForwardingRule deletes PortForwardingRule
func (cli *ZSClient) DeletePortForwardingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-forwarding/{uuid}", uuid, string(deleteMode))
}
