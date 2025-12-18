// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDatabaseBackup queries DatabaseBackup list
func (cli *ZSClient) QueryDatabaseBackup(params param.QueryParam) ([]view.DatabaseBackupInventoryView, error) {
	var resp []view.DatabaseBackupInventoryView
	return resp, cli.List("v1/database-backups", &params, &resp)
}
