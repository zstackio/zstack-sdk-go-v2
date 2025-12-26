// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddOssBucketFromRemote adds OssBucketFromRemote
func (cli *ZSClient) AddOssBucketFromRemote(params param.AddOssBucketFromRemoteParam) (*view.AddOssBucketFromRemoteEventView, error) {
	resp := view.AddOssBucketFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/oss-bucket", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
