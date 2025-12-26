// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddDisasterImageStoreBackupStorage adds DisasterImageStoreBackupStorage
func (cli *ZSClient) AddDisasterImageStoreBackupStorage(params param.AddDisasterImageStoreBackupStorageParam) (*view.AddImageStoreBackupStorageEventView, error) {
	resp := view.AddImageStoreBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/image-store/disaster", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
