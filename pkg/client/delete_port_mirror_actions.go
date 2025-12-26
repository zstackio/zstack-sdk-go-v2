// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePortMirror deletes PortMirror
func (cli *ZSClient) DeletePortMirror(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-mirrors/{uuid}", uuid, string(deleteMode))
}
