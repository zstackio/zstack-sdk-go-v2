// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVmBackup deletes VmBackup
func (cli *ZSClient) DeleteVmBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-backups/{groupUuid}", uuid, string(deleteMode))
}
