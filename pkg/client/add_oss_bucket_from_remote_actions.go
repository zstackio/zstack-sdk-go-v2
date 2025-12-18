// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddOssBucketFromRemote 操作AddOssBucketFromRemote
func (cli *ZSClient) AddOssBucketFromRemote(params param.AddOssBucketFromRemoteParam) (*view.AddOssBucketFromRemoteEventView, error) {
	resp := view.AddOssBucketFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/oss-bucket", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

