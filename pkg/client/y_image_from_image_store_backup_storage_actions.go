// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoveryImageFromImageStoreBackupStorage 操作yImageFromImageStoreBackupStorage
func (cli *ZSClient) RecoveryImageFromImageStoreBackupStorage(uuid string, params param.RecoveryImageFromImageStoreBackupStorageParam) (*view.RecoveryImageFromImageStoreBackupStorageEventView, error) {
	resp := view.RecoveryImageFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

