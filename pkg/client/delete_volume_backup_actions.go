// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteVolumeBackup deletes VolumeBackup
func (cli *ZSClient) DeleteVolumeBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-backups/{uuid}", uuid, string(deleteMode))
}
