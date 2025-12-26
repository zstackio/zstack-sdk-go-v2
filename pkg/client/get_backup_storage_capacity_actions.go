// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetBackupStorageCapacity gets BackupStorageCapacity by uuid
func (cli *ZSClient) GetBackupStorageCapacity(uuid string) (*view.GetBackupStorageCapacityView, error) {
	var resp view.GetBackupStorageCapacityView
	if err := cli.Get("v1/backup-storage/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
