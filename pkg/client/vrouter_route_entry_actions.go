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
func (cli *ZSClient) AddVRouterRouteEntry(routeTableUuid string, params param.AddVRouterRouteEntryParam) (*view.VRouterRouteEntryInventoryView, error) {
	resp := view.VRouterRouteEntryInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vrouter-route-tables/%s/route-entries", routeTableUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVRouterRouteEntry queries VRouterRouteEntry list
func (cli *ZSClient) QueryVRouterRouteEntry(params *param.QueryParam) ([]view.VRouterRouteEntryInventoryView, error) {
	var resp []view.VRouterRouteEntryInventoryView
	return resp, cli.List("v1/vrouter-route-tables/route-entries", params, &resp)
}

func (cli *ZSClient) GetVRouterRouteEntry(uuid string) (*view.VRouterRouteEntryInventoryView, error) {
	var resp view.VRouterRouteEntryInventoryView
	if err := cli.Get("v1/vrouter-route-tables/route-entries", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterRouteEntry Pagination
func (cli *ZSClient) PageVRouterRouteEntry(params *param.QueryParam) ([]view.VRouterRouteEntryInventoryView, int, error) {
	var vRouterRouteEntries []view.VRouterRouteEntryInventoryView
	total, err := cli.Page("v1/vrouter-route-tables/route-entries", params, &vRouterRouteEntries)
	return vRouterRouteEntries, total, err
}
