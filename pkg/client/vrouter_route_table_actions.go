// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVRouterRouteTable queries VRouterRouteTable list
func (cli *ZSClient) QueryVRouterRouteTable(params *param.QueryParam) ([]view.VRouterRouteTableInventoryView, error) {
	var resp []view.VRouterRouteTableInventoryView
	return resp, cli.List("v1/vrouter-route-tables", params, &resp)
}
// DeleteVRouterRouteTable deletes VRouterRouteTable
func (cli *ZSClient) DeleteVRouterRouteTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables/{uuid}", uuid, string(deleteMode))
}
// GetVRouterRouteTable gets VRouterRouteTable by uuid
func (cli *ZSClient) GetVRouterRouteTable(uuid string) (*view.VRouterRouteEntryAOView, error) {
	var resp view.VRouterRouteEntryAOView
	if err := cli.Get("v1/vrouter-route-tables/vrouter/{virtualRouterVmUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVRouterRouteTable creates VRouterRouteTable
func (cli *ZSClient) CreateVRouterRouteTable(params param.CreateVRouterRouteTableParam) (*view.VRouterRouteTableInventoryView, error) {
	var resp view.CreateVRouterRouteTableEventView
	if err := cli.Post("v1/vrouter-route-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateVRouterRouteTable updates VRouterRouteTable
func (cli *ZSClient) UpdateVRouterRouteTable(uuid string, params param.UpdateVRouterRouteTableParam) (*view.VRouterRouteTableInventoryView, error) {
	var resp view.UpdateVRouterRouteTableEventView
	if err := cli.Put("v1/vrouter-route-tables/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
