// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVRouterRouteEntry deletes VRouterRouteEntry
func (cli *ZSClient) DeleteVRouterRouteEntry(routeTableUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vrouter-route-tables", fmt.Sprintf("%s/route-entries/%s", routeTableUuid, uuid), string(deleteMode))
}
// AddVRouterRouteEntry adds VRouterRouteEntry
func (cli *ZSClient) AddVRouterRouteEntry(params param.AddVRouterRouteEntryParam) (*view.VRouterRouteEntryInventoryView, error) {
	var resp view.AddVRouterRouteEntryEventView
	if err := cli.Post("v1/vrouter-route-tables/{routeTableUuid}/route-entries", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
