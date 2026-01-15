// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterVm queries VirtualRouterVm list
func (cli *ZSClient) QueryVirtualRouterVm(params *param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List("v1/vm-instances/appliances/virtual-routers", params, &resp)
}

// PageVirtualRouterVm Pagination
func (cli *ZSClient) PageVirtualRouterVm(params *param.QueryParam) ([]view.ApplianceVmInventoryView, int, error) {
	var virtualRouterVms []view.ApplianceVmInventoryView
	total, err := cli.Page("v1/vm-instances/appliances/virtual-routers", params, &virtualRouterVms)
	return virtualRouterVms, total, err
}
