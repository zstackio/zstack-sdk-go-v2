// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncVmBackup operates on SyncVmBackup
func (cli *ZSClient) SyncVmBackup(uuid string, params param.SyncVmBackupParam) (*view.SyncVmBackupEventView, error) {
	resp := view.SyncVmBackupEventView{}
	if err := cli.Put("v1/vm-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
