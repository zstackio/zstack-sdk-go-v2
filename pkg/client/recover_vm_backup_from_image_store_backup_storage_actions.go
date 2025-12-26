// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverVmBackupFromImageStoreBackupStorage(uuid string, params param.RecoverVmBackupFromImageStoreBackupStorageParam) (*view.RecoverVmBackupFromImageStoreBackupStorageEventView, error) {
	resp := view.RecoverVmBackupFromImageStoreBackupStorageEventView{}
	if err := cli.Put("v1/vm-backups/{groupUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
