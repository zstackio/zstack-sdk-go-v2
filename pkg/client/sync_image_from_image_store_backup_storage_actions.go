// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncImageFromImageStoreBackupStorage 操作SyncImageFromImageStoreBackupStorage
func (cli *ZSClient) SyncImageFromImageStoreBackupStorage(uuid string, params param.SyncImageFromImageStoreBackupStorageParam) (*view.SyncImageFromImageStoreBackupStorageEventView, error) {
	resp := view.SyncImageFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

