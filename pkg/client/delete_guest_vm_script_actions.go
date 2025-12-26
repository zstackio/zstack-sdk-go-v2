// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteGuestVmScript deletes GuestVmScript
func (cli *ZSClient) DeleteGuestVmScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scripts/{uuid}", uuid, string(deleteMode))
}
