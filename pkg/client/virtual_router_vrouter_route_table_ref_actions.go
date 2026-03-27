// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterVRouterRouteTableRef queries VirtualRouterVRouterRouteTableRef list
func (cli *ZSClient) QueryVirtualRouterVRouterRouteTableRef(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterVRouterRouteTableRefInventoryView, error) {
	var resp []view.VirtualRouterVRouterRouteTableRefInventoryView
	return resp, cli.List(ctx, "v1/vrouter-route-tables/virtual-router-refs", params, &resp)
}

func (cli *ZSClient) GetVirtualRouterVRouterRouteTableRef(ctx context.Context, uuid string) (*view.VirtualRouterVRouterRouteTableRefInventoryView, error) {
	var resp view.VirtualRouterVRouterRouteTableRefInventoryView
	if err := cli.Get(ctx, "v1/vrouter-route-tables/virtual-router-refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVirtualRouterVRouterRouteTableRef Pagination
func (cli *ZSClient) PageVirtualRouterVRouterRouteTableRef(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterVRouterRouteTableRefInventoryView, int, error) {
	var virtualRouterVRouterRouteTableRefs []view.VirtualRouterVRouterRouteTableRefInventoryView
	total, err := cli.Page(ctx, "v1/vrouter-route-tables/virtual-router-refs", params, &virtualRouterVRouterRouteTableRefs)
	return virtualRouterVRouterRouteTableRefs, total, err
}
