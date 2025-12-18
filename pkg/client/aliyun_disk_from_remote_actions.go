// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunDiskFromRemote 创建AliyunDiskFromRemote
func (cli *ZSClient) CreateAliyunDiskFromRemote(params param.CreateAliyunDiskFromRemoteParam) (*view.CreateAliyunDiskFromRemoteEventView, error) {
	resp := view.CreateAliyunDiskFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/disk", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

