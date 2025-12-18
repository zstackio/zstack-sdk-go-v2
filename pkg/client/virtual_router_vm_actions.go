// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVirtualRouterVm 查询VirtualRouterVm列表
func (cli *ZSClient) QueryVirtualRouterVm(params param.QueryParam) ([]view.QueryApplianceVmView, error) {
	var resp []view.QueryApplianceVmView
	return resp, cli.List("v1/vm-instances/appliances/virtual-routers", &params, &resp)
}

