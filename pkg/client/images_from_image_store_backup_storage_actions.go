// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetImagesFromImageStoreBackupStorage 获取ImagesFromImageStoreBackupStorage详情
func (cli *ZSClient) GetImagesFromImageStoreBackupStorage(uuid string) (*view.GetImagesFromImageStoreBackupStorageView, error) {
	var resp view.GetImagesFromImageStoreBackupStorageView
	if err := cli.Get("v1/backup-storage/{uuid}/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

