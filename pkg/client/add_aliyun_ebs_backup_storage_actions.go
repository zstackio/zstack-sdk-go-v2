// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAliyunEbsBackupStorage adds AliyunEbsBackupStorage
func (cli *ZSClient) AddAliyunEbsBackupStorage(params param.AddAliyunEbsBackupStorageParam) (*view.AddBackupStorageEventView, error) {
	resp := view.AddBackupStorageEventView{}
	if err := cli.Post("v1/backup-storage/aliyun/ebs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
