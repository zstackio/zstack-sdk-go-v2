// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryExternalBackup 查询ExternalBackup列表
func (cli *ZSClient) QueryExternalBackup(params param.QueryParam) ([]view.QueryExternalBackupView, error) {
	var resp []view.QueryExternalBackupView
	return resp, cli.List("v1/externalbackup", &params, &resp)
}

