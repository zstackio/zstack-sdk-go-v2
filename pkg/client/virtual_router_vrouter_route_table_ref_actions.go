// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterVRouterRouteTableRef queries VirtualRouterVRouterRouteTableRef list
func (cli *ZSClient) QueryVirtualRouterVRouterRouteTableRef(params *param.QueryParam) ([]view.VirtualRouterVRouterRouteTableRefInventoryView, error) {
	var resp []view.VirtualRouterVRouterRouteTableRefInventoryView
	return resp, cli.List("v1/vrouter-route-tables/virtual-router-refs", params, &resp)
}

// PageVirtualRouterVRouterRouteTableRef Pagination
func (cli *ZSClient) PageVirtualRouterVRouterRouteTableRef(params *param.QueryParam) ([]view.VirtualRouterVRouterRouteTableRefInventoryView, int, error) {
	var virtualRouterVRouterRouteTableRefs []view.VirtualRouterVRouterRouteTableRefInventoryView
	total, err := cli.Page("v1/vrouter-route-tables/virtual-router-refs", params, &virtualRouterVRouterRouteTableRefs)
	return virtualRouterVRouterRouteTableRefs, total, err
}
