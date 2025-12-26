// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetBackupStorageTypes gets BackupStorageTypes by uuid
func (cli *ZSClient) GetBackupStorageTypes(uuid string) (*view.GetBackupStorageTypesView, error) {
	var resp view.GetBackupStorageTypesView
	if err := cli.Get("v1/backup-storage/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
