// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunDiskFromRemote creates AliyunDiskFromRemote
func (cli *ZSClient) CreateAliyunDiskFromRemote(params param.CreateAliyunDiskFromRemoteParam) (*view.CreateAliyunDiskFromRemoteEventView, error) {
	resp := view.CreateAliyunDiskFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
