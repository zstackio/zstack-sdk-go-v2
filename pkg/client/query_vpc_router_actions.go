// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVpcRouter queries VpcRouter list
func (cli *ZSClient) QueryVpcRouter(params *param.QueryParam) ([]view.VpcRouterVmInventoryView, error) {
	var resp []view.VpcRouterVmInventoryView
	return resp, cli.List("v1/vpc/virtual-routers", params, &resp)
}
