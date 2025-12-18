// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryZBoxBackup 查询ZBoxBackup列表
func (cli *ZSClient) QueryZBoxBackup(params param.QueryParam) ([]view.QueryZBoxBackupView, error) {
	var resp []view.QueryZBoxBackupView
	return resp, cli.List("v1/externalbackup/zbox", &params, &resp)
}

