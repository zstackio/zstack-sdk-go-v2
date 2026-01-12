// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterOffering queries VirtualRouterOffering list
func (cli *ZSClient) QueryVirtualRouterOffering(params *param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, error) {
	var resp []view.VirtualRouterOfferingInventoryView
	return resp, cli.List("v1/instance-offerings/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVirtualRouterOffering(uuid string) (*view.VirtualRouterOfferingInventoryView, error) {
	var resp view.VirtualRouterOfferingInventoryView
	if err := cli.Get("v1/instance-offerings/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	if err := cli.Put("v1/instance-offerings/virtual-routers", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
