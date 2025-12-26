// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetDatabaseBackupFromImageStore gets DatabaseBackupFromImageStore by uuid
func (cli *ZSClient) GetDatabaseBackupFromImageStore(uuid string) (*view.GetDatabaseBackupFromImageStoreView, error) {
	var resp view.GetDatabaseBackupFromImageStoreView
	if err := cli.Get("v1/database-backups/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
