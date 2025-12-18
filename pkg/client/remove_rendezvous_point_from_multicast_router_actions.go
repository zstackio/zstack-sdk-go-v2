// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveRendezvousPointFromMulticastRouter 操作RemoveRendezvousPointFromMulticastRouter
func (cli *ZSClient) RemoveRendezvousPointFromMulticastRouter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/multicast/virtual-routers/{uuid}/RendezvousPoint", uuid, string(deleteMode))
}

