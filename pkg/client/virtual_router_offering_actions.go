// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterOffering queries VirtualRouterOffering list
func (cli *ZSClient) QueryVirtualRouterOffering(params *param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, error) {
	var resp []view.VirtualRouterOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/virtual-routers", params, &resp)
}
// CreateVirtualRouterOffering creates VirtualRouterOffering
func (cli *ZSClient) CreateVirtualRouterOffering(params param.CreateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.CreateInstanceOfferingEventView
	if err := cli.Post("v1/instance-offerings/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateVirtualRouterOffering updates VirtualRouterOffering
func (cli *ZSClient) UpdateVirtualRouterOffering(uuid string, params param.UpdateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	var resp view.UpdateInstanceOfferingEventView
	if err := cli.Put("v1/instance-offerings/virtual-routers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
