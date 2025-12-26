// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncDatabaseBackup operates on SyncDatabaseBackup
func (cli *ZSClient) SyncDatabaseBackup(uuid string, params param.SyncDatabaseBackupParam) (*view.SyncDatabaseBackupEventView, error) {
	resp := view.SyncDatabaseBackupEventView{}
	if err := cli.Put("v1/database-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
