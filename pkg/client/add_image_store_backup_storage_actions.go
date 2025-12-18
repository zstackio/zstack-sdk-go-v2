// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddImageStoreBackupStorage 操作AddImageStoreBackupStorage
func (cli *ZSClient) AddImageStoreBackupStorage(params param.AddImageStoreBackupStorageParam) (*view.AddImageStoreBackupStorageEventView, error) {
	resp := view.AddImageStoreBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/image-store", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

