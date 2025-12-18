// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncBackupFromImageStoreBackupStorage 操作SyncBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncBackupFromImageStoreBackupStorage(uuid string, params param.SyncBackupFromImageStoreBackupStorageParam) (*view.SyncBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.SyncBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

