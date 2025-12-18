// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncDatabaseBackupFromImageStoreBackupStorage 操作SyncDatabaseBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncDatabaseBackupFromImageStoreBackupStorage(uuid string, params param.SyncDatabaseBackupFromImageStoreBackupStorageParam) (*view.SyncDatabaseBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.SyncDatabaseBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/database-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

