// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunVpcVirtualRouterEntryRemote creates AliyunVpcVirtualRouterEntryRemote
func (cli *ZSClient) CreateAliyunVpcVirtualRouterEntryRemote(params param.CreateAliyunVpcVirtualRouterEntryRemoteParam) (*view.CreateAliyunVpcVirtualRouterEntryRemoteEventView, error) {
	resp := view.CreateAliyunVpcVirtualRouterEntryRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/route-entry", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
