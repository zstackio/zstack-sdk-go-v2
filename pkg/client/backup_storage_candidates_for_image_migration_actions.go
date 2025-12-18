// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBackupStorageCandidatesForImageMigration 获取BackupStorageCandidatesForImageMigration详情
func (cli *ZSClient) GetBackupStorageCandidatesForImageMigration(uuid string) (*view.GetBackupStorageCandidatesForImageMigrationView, error) {
	var resp view.GetBackupStorageCandidatesForImageMigrationView
	if err := cli.Get("v1/backup-storage/{srcBackupStorageUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

