// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVRouterRouteEntry deletes VRouterRouteEntry
func (cli *ZSClient) DeleteVRouterRouteEntry(routeTableUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vrouter-route-tables", routeTableUuid, fmt.Sprintf("route-entries/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddVRouterRouteEntry adds VRouterRouteEntry
func (cli *ZSClient) AddVRouterRouteEntry(params param.AddVRouterRouteEntryParam) (*view.VRouterRouteEntryInventoryView, error) {
	resp := view.VRouterRouteEntryInventoryView{}
	if err := cli.Post("v1/vrouter-route-tables/{routeTableUuid}/route-entries", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVRouterRouteEntry queries VRouterRouteEntry list
func (cli *ZSClient) QueryVRouterRouteEntry(params *param.QueryParam) ([]view.VRouterRouteEntryInventoryView, error) {
	var resp []view.VRouterRouteEntryInventoryView
	return resp, cli.List("v1/vrouter-route-tables/route-entries", params, &resp)
}

// PageVRouterRouteEntry Pagination
func (cli *ZSClient) PageVRouterRouteEntry(params *param.QueryParam) ([]view.VRouterRouteEntryInventoryView, int, error) {
	var vRouterRouteEntries []view.VRouterRouteEntryInventoryView
	total, err := cli.Page("v1/vrouter-route-tables/route-entries", params, &vRouterRouteEntries)
	return vRouterRouteEntries, total, err
}
