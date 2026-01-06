// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterVm queries VirtualRouterVm list
func (cli *ZSClient) QueryVirtualRouterVm(params *param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List("v1/vm-instances/appliances/virtual-routers", params, &resp)
}
