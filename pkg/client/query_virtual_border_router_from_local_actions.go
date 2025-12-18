// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVirtualBorderRouterFromLocal queries VirtualBorderRouterFromLocal list
func (cli *ZSClient) QueryVirtualBorderRouterFromLocal(params param.QueryParam) ([]view.VirtualBorderRouterInventoryView, error) {
	var resp []view.VirtualBorderRouterInventoryView
	return resp, cli.List("v1/hybrid/aliyun/border-router", &params, &resp)
}
