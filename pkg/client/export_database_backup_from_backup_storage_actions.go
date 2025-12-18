// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExportDatabaseBackupFromBackupStorage 操作ExportDatabaseBackupFromBackupStorage
func (cli *ZSClient) ExportDatabaseBackupFromBackupStorage(uuid string, params param.ExportDatabaseBackupFromBackupStorageParam) (*view.ExportDatabaseBackupFromBackupStorageEventView, error) {
	resp := view.ExportDatabaseBackupFromBackupStorageEventView{}
	if err := cli.Put("v1/database-backups/{databaseBackupUuid}/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

