// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBackupStorageTypes 获取BackupStorageTypes详情
func (cli *ZSClient) GetBackupStorageTypes(uuid string) (*view.GetBackupStorageTypesView, error) {
	var resp view.GetBackupStorageTypesView
	if err := cli.Get("v1/backup-storage/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

