// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryDatabaseBackup queries DatabaseBackup list
func (cli *ZSClient) QueryDatabaseBackup(params *param.QueryParam) ([]view.DatabaseBackupInventoryView, error) {
	var resp []view.DatabaseBackupInventoryView
	return resp, cli.List("v1/database-backups", params, &resp)
}
