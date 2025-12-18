// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncDatabaseBackup operates on SyncDatabaseBackup
func (cli *ZSClient) SyncDatabaseBackup(uuid string, params param.SyncDatabaseBackupParam) (*view.SyncDatabaseBackupEventView, error) {
	resp := view.SyncDatabaseBackupEventView{}
	if err := cli.Put("v1/database-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
