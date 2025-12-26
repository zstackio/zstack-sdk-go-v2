// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverBackupFromImageStoreBackupStorage(uuid string, params param.RecoverBackupFromImageStoreBackupStorageParam) (*view.RecoverBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.RecoverBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
