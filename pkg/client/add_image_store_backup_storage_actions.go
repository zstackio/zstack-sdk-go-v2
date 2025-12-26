// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddImageStoreBackupStorage adds ImageStoreBackupStorage
func (cli *ZSClient) AddImageStoreBackupStorage(params param.AddImageStoreBackupStorageParam) (*view.AddImageStoreBackupStorageEventView, error) {
	resp := view.AddImageStoreBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/image-store", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
