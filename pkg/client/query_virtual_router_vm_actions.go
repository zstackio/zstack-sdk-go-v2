// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVirtualRouterVm queries VirtualRouterVm list
func (cli *ZSClient) QueryVirtualRouterVm(params param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List("v1/vm-instances/appliances/virtual-routers", &params, &resp)
}
