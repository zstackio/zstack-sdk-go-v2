// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetOssBackupBucketFromRemote gets OssBackupBucketFromRemote by uuid
func (cli *ZSClient) GetOssBackupBucketFromRemote(uuid string) (*view.GetOssBackupBucketFromRemoteView, error) {
	var resp view.GetOssBackupBucketFromRemoteView
	if err := cli.Get("v1/hybrid/backup-mysql/oss", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
