// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmConsolePassword deletes VmConsolePassword
func (cli *ZSClient) DeleteVmConsolePassword(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/console-password", uuid, string(deleteMode))
}
