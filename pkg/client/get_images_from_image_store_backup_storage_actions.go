// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetImagesFromImageStoreBackupStorage gets ImagesFromImageStoreBackupStorage by uuid
func (cli *ZSClient) GetImagesFromImageStoreBackupStorage(uuid string) (*view.GetImagesFromImageStoreBackupStorageView, error) {
	var resp view.GetImagesFromImageStoreBackupStorageView
	if err := cli.Get("v1/backup-storage/{uuid}/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
