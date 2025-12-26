// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAliyunVirtualRouterFromLocal queries AliyunVirtualRouterFromLocal list
func (cli *ZSClient) QueryAliyunVirtualRouterFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouterInventoryView, error) {
	var resp []view.VpcVirtualRouterInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vrouter", params, &resp)
}
