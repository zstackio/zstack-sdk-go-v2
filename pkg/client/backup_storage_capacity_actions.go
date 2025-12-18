// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBackupStorageCapacity 获取BackupStorageCapacity详情
func (cli *ZSClient) GetBackupStorageCapacity(uuid string) (*view.GetBackupStorageCapacityView, error) {
	var resp view.GetBackupStorageCapacityView
	if err := cli.Get("v1/backup-storage/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

