// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateOssBackupBucketRemote creates OssBackupBucketRemote
func (cli *ZSClient) CreateOssBackupBucketRemote(params param.CreateOssBackupBucketRemoteParam) (*view.CreateOssBackupBucketRemoteEventView, error) {
	resp := view.CreateOssBackupBucketRemoteEventView{}
	if err := cli.Post("v1/hybrid/backup-mysql/oss", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
