// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVirtualBorderRouterFromLocal queries VirtualBorderRouterFromLocal list
func (cli *ZSClient) QueryVirtualBorderRouterFromLocal(params *param.QueryParam) ([]view.VirtualBorderRouterInventoryView, error) {
	var resp []view.VirtualBorderRouterInventoryView
	return resp, cli.List("v1/hybrid/aliyun/border-router", params, &resp)
}
