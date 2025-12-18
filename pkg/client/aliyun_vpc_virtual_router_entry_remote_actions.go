// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunVpcVirtualRouterEntryRemote 创建AliyunVpcVirtualRouterEntryRemote
func (cli *ZSClient) CreateAliyunVpcVirtualRouterEntryRemote(params param.CreateAliyunVpcVirtualRouterEntryRemoteParam) (*view.CreateAliyunVpcVirtualRouterEntryRemoteEventView, error) {
	resp := view.CreateAliyunVpcVirtualRouterEntryRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/route-entry", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

