// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncBackupFromImageStoreBackupStorage operates on SyncBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncBackupFromImageStoreBackupStorage(uuid string, params param.SyncBackupFromImageStoreBackupStorageParam) (*view.SyncBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.SyncBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
