// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeBackupStorageState changes BackupStorageState
func (cli *ZSClient) ChangeBackupStorageState(uuid string, params param.ChangeBackupStorageStateParam) (*view.ChangeBackupStorageStateEventView, error) {
	resp := view.ChangeBackupStorageStateEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
