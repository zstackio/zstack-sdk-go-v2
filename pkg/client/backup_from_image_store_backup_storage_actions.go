// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoverBackupFromImageStoreBackupStorage 操作BackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverBackupFromImageStoreBackupStorage(uuid string, params param.RecoverBackupFromImageStoreBackupStorageParam) (*view.RecoverBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.RecoverBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/volume-backups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

