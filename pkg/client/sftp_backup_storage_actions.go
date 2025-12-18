// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySftpBackupStorage 查询SftpBackupStorage列表
func (cli *ZSClient) QuerySftpBackupStorage(params param.QueryParam) ([]view.QuerySftpBackupStorageView, error) {
	var resp []view.QuerySftpBackupStorageView
	return resp, cli.List("v1/backup-storage/sftp", &params, &resp)
}

