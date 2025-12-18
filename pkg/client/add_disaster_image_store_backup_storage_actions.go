// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddDisasterImageStoreBackupStorage 操作AddDisasterImageStoreBackupStorage
func (cli *ZSClient) AddDisasterImageStoreBackupStorage(params param.AddDisasterImageStoreBackupStorageParam) (*view.AddImageStoreBackupStorageEventView, error) {
	resp := view.AddImageStoreBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/image-store/disaster", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

