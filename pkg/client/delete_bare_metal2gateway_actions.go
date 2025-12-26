// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBareMetal2Gateway deletes BareMetal2Gateway
func (cli *ZSClient) DeleteBareMetal2Gateway(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/gateways/{uuid}", uuid, string(deleteMode))
}
