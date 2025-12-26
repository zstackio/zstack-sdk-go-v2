// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMulticastRouter deletes MulticastRouter
func (cli *ZSClient) DeleteMulticastRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/multicast/virtual-routers/{uuid}", uuid, string(deleteMode))
}
