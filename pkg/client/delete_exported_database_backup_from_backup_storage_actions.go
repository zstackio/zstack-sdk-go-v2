// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteExportedDatabaseBackupFromBackupStorage deletes ExportedDatabaseBackupFromBackupStorage
func (cli *ZSClient) DeleteExportedDatabaseBackupFromBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/exported-database-backup/{databaseBackupUuid}/backup-storage/{backupStorageUuid}", uuid, string(deleteMode))
}
