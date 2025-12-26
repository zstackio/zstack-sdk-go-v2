// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetOssBucketFileFromRemote gets OssBucketFileFromRemote by uuid
func (cli *ZSClient) GetOssBucketFileFromRemote(uuid string) (*view.GetOssBucketFileFromRemoteView, error) {
	var resp view.GetOssBucketFileFromRemoteView
	if err := cli.Get("v1/hybrid/oss/file/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
