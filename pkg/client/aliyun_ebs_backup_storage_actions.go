// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunEbsBackupStorage 查询AliyunEbsBackupStorage列表
func (cli *ZSClient) QueryAliyunEbsBackupStorage(params param.QueryParam) ([]view.QueryBackupStorageView, error) {
	var resp []view.QueryBackupStorageView
	return resp, cli.List("v1/backup-storage/aliyun/ebs", &params, &resp)
}

