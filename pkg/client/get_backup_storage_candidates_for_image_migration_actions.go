// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetBackupStorageCandidatesForImageMigration gets BackupStorageCandidatesForImageMigration by uuid
func (cli *ZSClient) GetBackupStorageCandidatesForImageMigration(uuid string) (*view.GetBackupStorageCandidatesForImageMigrationView, error) {
	var resp view.GetBackupStorageCandidatesForImageMigrationView
	if err := cli.Get("v1/backup-storage/{srcBackupStorageUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
