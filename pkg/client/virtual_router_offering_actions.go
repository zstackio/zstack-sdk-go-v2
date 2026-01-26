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

// PageVirtualRouterOffering Pagination
func (cli *ZSClient) PageVirtualRouterOffering(params *param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, int, error) {
	var virtualRouterOfferings []view.VirtualRouterOfferingInventoryView
	total, err := cli.Page("v1/instance-offerings/virtual-routers", params, &virtualRouterOfferings)
	return virtualRouterOfferings, total, err
}
// CreateVirtualRouterOffering creates VirtualRouterOffering
func (cli *ZSClient) CreateVirtualRouterOffering(params param.CreateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post("v1/instance-offerings/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVirtualRouterOffering updates VirtualRouterOffering
func (cli *ZSClient) UpdateVirtualRouterOffering(uuid string, params param.UpdateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithRespKey("v1/instance-offerings/virtual-routers", uuid, "", map[string]interface{}{
		"updateVirtualRouterOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
