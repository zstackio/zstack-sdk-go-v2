// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoveryImageFromImageStoreBackupStorage operates on yImageFromImageStoreBackupStorage
func (cli *ZSClient) RecoveryImageFromImageStoreBackupStorage(uuid string, params param.RecoveryImageFromImageStoreBackupStorageParam) (*view.RecoveryImageFromImageStoreBackupStorageEventView, error) {
	resp := view.RecoveryImageFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
