// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVRouterRouteEntry deletes VRouterRouteEntry
func (cli *ZSClient) DeleteVRouterRouteEntry(ctx context.Context, routeTableUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/vrouter-route-tables", routeTableUuid, fmt.Sprintf("route-entries/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddVRouterRouteEntry adds VRouterRouteEntry
func (cli *ZSClient) AddVRouterRouteEntry(ctx context.Context, params param.AddVRouterRouteEntryParam) (*view.VRouterRouteEntryInventoryView, error) {
	resp := view.VRouterRouteEntryInventoryView{}
	if err := cli.Post(ctx, "v1/vrouter-route-tables/{routeTableUuid}/route-entries", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVRouterRouteEntry queries VRouterRouteEntry list
func (cli *ZSClient) QueryVRouterRouteEntry(ctx context.Context, params *param.QueryParam) ([]view.VRouterRouteEntryInventoryView, error) {
	var resp []view.VRouterRouteEntryInventoryView
	return resp, cli.List(ctx, "v1/vrouter-route-tables/route-entries", params, &resp)
}

func (cli *ZSClient) GetVRouterRouteEntry(ctx context.Context, uuid string) (*view.VRouterRouteEntryInventoryView, error) {
	var resp view.VRouterRouteEntryInventoryView
	if err := cli.Get(ctx, "v1/vrouter-route-tables/route-entries", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterRouteEntry Pagination
func (cli *ZSClient) PageVRouterRouteEntry(ctx context.Context, params *param.QueryParam) ([]view.VRouterRouteEntryInventoryView, int, error) {
	var vRouterRouteEntries []view.VRouterRouteEntryInventoryView
	total, err := cli.Page(ctx, "v1/vrouter-route-tables/route-entries", params, &vRouterRouteEntries)
	return vRouterRouteEntries, total, err
}
