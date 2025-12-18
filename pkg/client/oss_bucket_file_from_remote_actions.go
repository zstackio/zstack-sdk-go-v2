// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetOssBucketFileFromRemote 获取OssBucketFileFromRemote详情
func (cli *ZSClient) GetOssBucketFileFromRemote(uuid string) (*view.GetOssBucketFileFromRemoteView, error) {
	var resp view.GetOssBucketFileFromRemoteView
	if err := cli.Get("v1/hybrid/oss/file/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

