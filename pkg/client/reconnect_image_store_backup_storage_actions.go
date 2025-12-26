// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectImageStoreBackupStorage operates on ReconnectImageStoreBackupStorage
func (cli *ZSClient) ReconnectImageStoreBackupStorage(uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ReconnectImageStoreBackupStorageEventView, error) {
	resp := view.ReconnectImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
