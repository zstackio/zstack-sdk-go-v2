// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunEbsBackupStorage 操作AddAliyunEbsBackupStorage
func (cli *ZSClient) AddAliyunEbsBackupStorage(params param.AddAliyunEbsBackupStorageParam) (*view.AddBackupStorageEventView, error) {
	resp := view.AddBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/aliyun/ebs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

