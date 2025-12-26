// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExportDatabaseBackupFromBackupStorage operates on ExportDatabaseBackupFromBackupStorage
func (cli *ZSClient) ExportDatabaseBackupFromBackupStorage(uuid string, params param.ExportDatabaseBackupFromBackupStorageParam) (*view.ExportDatabaseBackupFromBackupStorageEventView, error) {
	resp := view.ExportDatabaseBackupFromBackupStorageEventView{}
	if err := cli.Put("v1/database-backups/{databaseBackupUuid}/backup-storage/{backupStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
