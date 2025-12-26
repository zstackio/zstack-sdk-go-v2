// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncImageFromImageStoreBackupStorage operates on SyncImageFromImageStoreBackupStorage
func (cli *ZSClient) SyncImageFromImageStoreBackupStorage(uuid string, params param.SyncImageFromImageStoreBackupStorageParam) (*view.SyncImageFromImageStoreBackupStorageEventView, error) {
	resp := view.SyncImageFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
