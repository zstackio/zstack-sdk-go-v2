// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncVmBackupFromImageStoreBackupStorage operates on SyncVmBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncVmBackupFromImageStoreBackupStorage(uuid string, params param.SyncVmBackupFromImageStoreBackupStorageParam) (*view.SyncVmBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.SyncVmBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
