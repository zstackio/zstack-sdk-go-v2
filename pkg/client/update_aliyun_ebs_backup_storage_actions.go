// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunEbsBackupStorage updates AliyunEbsBackupStorage
func (cli *ZSClient) UpdateAliyunEbsBackupStorage(uuid string, params param.UpdateAliyunEbsBackupStorageParam) (*view.UpdateBackupStorageEventView, error) {
	resp := view.UpdateBackupStorageEventView{}
	if err := cli.Put("v1/backup-storage/aliyun/ebs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
