// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteBackupStorage deletes BackupStorage
func (cli *ZSClient) DeleteBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/{uuid}", uuid, string(deleteMode))
}
