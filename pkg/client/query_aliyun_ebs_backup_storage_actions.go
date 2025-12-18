// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunEbsBackupStorage queries AliyunEbsBackupStorage list
func (cli *ZSClient) QueryAliyunEbsBackupStorage(params param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/aliyun/ebs", &params, &resp)
}
