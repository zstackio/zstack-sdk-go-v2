// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteExportedDatabaseBackupFromBackupStorage 删除ExportedDatabaseBackupFromBackupStorage
func (cli *ZSClient) DeleteExportedDatabaseBackupFromBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/exported-database-backup/{databaseBackupUuid}/backup-storage/{backupStorageUuid}", uuid, string(deleteMode))
}

