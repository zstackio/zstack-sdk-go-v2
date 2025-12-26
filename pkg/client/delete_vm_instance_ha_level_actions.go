// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmInstanceHaLevel deletes VmInstanceHaLevel
func (cli *ZSClient) DeleteVmInstanceHaLevel(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{uuid}/ha-levels", uuid, string(deleteMode))
}
