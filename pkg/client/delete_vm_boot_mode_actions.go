// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmBootMode deletes VmBootMode
func (cli *ZSClient) DeleteVmBootMode(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/bootmode", uuid, string(deleteMode))
}
