// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePortMirrorSession deletes PortMirrorSession
func (cli *ZSClient) DeletePortMirrorSession(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-mirrors/sessons/{uuid}", uuid, string(deleteMode))
}
