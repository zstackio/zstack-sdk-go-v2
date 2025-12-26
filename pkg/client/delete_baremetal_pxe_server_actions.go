// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBaremetalPxeServer deletes BaremetalPxeServer
func (cli *ZSClient) DeleteBaremetalPxeServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/pxeservers/{uuid}", uuid, string(deleteMode))
}
