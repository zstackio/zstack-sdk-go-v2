// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenterBackupStorage 查询VCenterBackupStorage列表
func (cli *ZSClient) QueryVCenterBackupStorage(params param.QueryParam) ([]view.QueryVCenterBackupStorageView, error) {
	var resp []view.QueryVCenterBackupStorageView
	return resp, cli.List("v1/vcenters/backup-storage", &params, &resp)
}

