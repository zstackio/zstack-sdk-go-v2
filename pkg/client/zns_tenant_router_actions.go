// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryZnsTenantRouter queries ZnsTenantRouter list
func (cli *ZSClient) QueryZnsTenantRouter(ctx context.Context, params *param.QueryParam) ([]view.ZnsTenantRouterInventoryView, error) {
	var resp []view.ZnsTenantRouterInventoryView
	return resp, cli.List(ctx, "v1/sdn-controller/zns/tenant-routers", params, &resp)
}

func (cli *ZSClient) GetZnsTenantRouter(ctx context.Context, uuid string) (*view.ZnsTenantRouterInventoryView, error) {
	var resp view.ZnsTenantRouterInventoryView
	if err := cli.Get(ctx, "v1/sdn-controller/zns/tenant-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZnsTenantRouter Pagination
func (cli *ZSClient) PageZnsTenantRouter(ctx context.Context, params *param.QueryParam) ([]view.ZnsTenantRouterInventoryView, int, error) {
	var znsTenantRouters []view.ZnsTenantRouterInventoryView
	total, err := cli.Page(ctx, "v1/sdn-controller/zns/tenant-routers", params, &znsTenantRouters)
	return znsTenantRouters, total, err
}
