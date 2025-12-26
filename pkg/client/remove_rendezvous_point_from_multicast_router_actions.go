// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveRendezvousPointFromMulticastRouter removes RendezvousPointFromMulticastRouter
func (cli *ZSClient) RemoveRendezvousPointFromMulticastRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/multicast/virtual-routers/{uuid}/RendezvousPoint", uuid, string(deleteMode))
}
