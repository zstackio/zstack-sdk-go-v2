// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVRouterRouteTable queries VRouterRouteTable list
func (cli *ZSClient) QueryVRouterRouteTable(params *param.QueryParam) ([]view.VRouterRouteTableInventoryView, error) {
	var resp []view.VRouterRouteTableInventoryView
	return resp, cli.List("v1/vrouter-route-tables", params, &resp)
}

func (cli *ZSClient) GetVRouterRouteTable(uuid string) (*view.VRouterRouteTableInventoryView, error) {
	var resp view.VRouterRouteTableInventoryView
	if err := cli.Get("v1/vrouter-route-tables", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVRouterRouteTable deletes VRouterRouteTable
func (cli *ZSClient) DeleteVRouterRouteTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables", uuid, string(deleteMode))
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
	if err := cli.Put("v1/vrouter-route-tables", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
