// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVirtualRouterOffering queries VirtualRouterOffering list
func (cli *ZSClient) QueryVirtualRouterOffering(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, error) {
	var resp []view.VirtualRouterOfferingInventoryView
	return resp, cli.List(ctx, "v1/instance-offerings/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVirtualRouterOffering(ctx context.Context, uuid string) (*view.VirtualRouterOfferingInventoryView, error) {
	var resp view.VirtualRouterOfferingInventoryView
	if err := cli.Get(ctx, "v1/instance-offerings/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVirtualRouterOffering Pagination
func (cli *ZSClient) PageVirtualRouterOffering(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, int, error) {
	var virtualRouterOfferings []view.VirtualRouterOfferingInventoryView
	total, err := cli.Page(ctx, "v1/instance-offerings/virtual-routers", params, &virtualRouterOfferings)
	return virtualRouterOfferings, total, err
}
// CreateVirtualRouterOffering creates VirtualRouterOffering
func (cli *ZSClient) CreateVirtualRouterOffering(ctx context.Context, params param.CreateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post(ctx, "v1/instance-offerings/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVirtualRouterOffering updates VirtualRouterOffering
func (cli *ZSClient) UpdateVirtualRouterOffering(ctx context.Context, uuid string, params param.UpdateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/instance-offerings/virtual-routers", uuid, "", map[string]interface{}{
		"updateVirtualRouterOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
