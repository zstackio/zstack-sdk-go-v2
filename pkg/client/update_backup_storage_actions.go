// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBackupStorage updates BackupStorage
func (cli *ZSClient) UpdateBackupStorage(uuid string, params param.UpdateBackupStorageParam) (*view.UpdateBackupStorageEventView, error) {
	resp := view.UpdateBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
