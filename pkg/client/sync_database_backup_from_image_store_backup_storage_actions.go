// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncDatabaseBackupFromImageStoreBackupStorage operates on SyncDatabaseBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncDatabaseBackupFromImageStoreBackupStorage(uuid string, params param.SyncDatabaseBackupFromImageStoreBackupStorageParam) (*view.SyncDatabaseBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.SyncDatabaseBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/database-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
