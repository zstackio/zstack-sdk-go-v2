// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateImageStoreBackupStorage updates ImageStoreBackupStorage
func (cli *ZSClient) UpdateImageStoreBackupStorage(uuid string, params param.UpdateImageStoreBackupStorageParam) (*view.UpdateBackupStorageEventView, error) {
	resp := view.UpdateBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
