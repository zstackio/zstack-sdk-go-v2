// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDatabaseBackup 查询DatabaseBackup列表
func (cli *ZSClient) QueryDatabaseBackup(params param.QueryParam) ([]view.QueryDatabaseBackupView, error) {
	var resp []view.QueryDatabaseBackupView
	return resp, cli.List("v1/database-backups", &params, &resp)
}

