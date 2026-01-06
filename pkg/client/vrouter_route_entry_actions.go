// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVRouterRouteEntry deletes VRouterRouteEntry
func (cli *ZSClient) DeleteVRouterRouteEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables/{routeTableUuid}/route-entries/{uuid}", uuid, string(deleteMode))
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
