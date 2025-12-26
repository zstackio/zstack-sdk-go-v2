// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteExportedImageFromBackupStorage deletes ExportedImageFromBackupStorage
func (cli *ZSClient) DeleteExportedImageFromBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/{backupStorageUuid}/exported-images/{imageUuid}", uuid, string(deleteMode))
}
