// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

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

// PageVRouterRouteTable Pagination
func (cli *ZSClient) PageVRouterRouteTable(params *param.QueryParam) ([]view.VRouterRouteTableInventoryView, int, error) {
	var vRouterRouteTables []view.VRouterRouteTableInventoryView
	total, err := cli.Page("v1/vrouter-route-tables", params, &vRouterRouteTables)
	return vRouterRouteTables, total, err
}
// DeleteVRouterRouteTable deletes VRouterRouteTable
func (cli *ZSClient) DeleteVRouterRouteTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vrouter-route-tables", uuid, string(deleteMode))
}
// CreateVRouterRouteTable creates VRouterRouteTable
func (cli *ZSClient) CreateVRouterRouteTable(params param.CreateVRouterRouteTableParam) (*view.VRouterRouteTableInventoryView, error) {
	resp := view.VRouterRouteTableInventoryView{}
	if err := cli.Post("v1/vrouter-route-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVRouterRouteTable updates VRouterRouteTable
func (cli *ZSClient) UpdateVRouterRouteTable(uuid string, params param.UpdateVRouterRouteTableParam) (*view.VRouterRouteTableInventoryView, error) {
	resp := view.VRouterRouteTableInventoryView{}
	if err := cli.PutWithRespKey("v1/vrouter-route-tables", uuid, "", map[string]interface{}{
		"updateVRouterRouteTable": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRouteTableVpcVRouterCandidate gets RouteTableVpcVRouterCandidate by uuid
func (cli *ZSClient) GetRouteTableVpcVRouterCandidate() (*view.VpcRouterVmInventoryView, error) {
	var resp view.VpcRouterVmInventoryView
	if err := cli.GetWithRespKey("v1/vpc/virtual-routers/get-vpc-candidate", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVRouterRouteTableFromVRouter operates on VRouterRouteTableFromVRouter
func (cli *ZSClient) DetachVRouterRouteTableFromVRouter(routeTableUuid string, virtualRouterVmUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vrouter-route-tables", routeTableUuid, fmt.Sprintf("detach/%s", virtualRouterVmUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AttachVRouterRouteTableToVRouter operates on VRouterRouteTableToVRouter
func (cli *ZSClient) AttachVRouterRouteTableToVRouter(routeTableUuid string, params param.AttachVRouterRouteTableToVRouterParam) (*view.VRouterRouteTableInventoryView, error) {
	resp := view.VRouterRouteTableInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vrouter-route-tables/%s/attach", routeTableUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
