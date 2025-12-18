// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateOssBucketRemote creates OssBucketRemote
func (cli *ZSClient) CreateOssBucketRemote(params param.CreateOssBucketRemoteParam) (*view.CreateOssBucketRemoteEventView, error) {
	resp := view.CreateOssBucketRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/oss-bucket/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
