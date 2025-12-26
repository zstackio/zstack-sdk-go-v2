// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryImageStoreBackupStorage queries ImageStoreBackupStorage list
func (cli *ZSClient) QueryImageStoreBackupStorage(params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, error) {
	var resp []view.ImageStoreBackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/image-store", params, &resp)
}
