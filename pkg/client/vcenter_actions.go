// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddVCenter adds VCenter
func (cli *ZSClient) AddVCenter(ctx context.Context, params param.AddVCenterParam) (*view.VCenterInventoryView, error) {
	resp := view.VCenterInventoryView{}
	if err := cli.Post(ctx, "v1/vcenters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// SyncVCenter operates on VCenter
func (cli *ZSClient) SyncVCenter(ctx context.Context, uuid string, params param.SyncVCenterParam) (*view.VCenterInventoryView, error) {
	resp := view.VCenterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vcenters", uuid, "", map[string]interface{}{
		"syncVCenter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVCenter queries VCenter list
func (cli *ZSClient) QueryVCenter(ctx context.Context, params *param.QueryParam) ([]view.VCenterInventoryView, error) {
	var resp []view.VCenterInventoryView
	return resp, cli.List(ctx, "v1/vcenters", params, &resp)
}

func (cli *ZSClient) GetVCenter(ctx context.Context, uuid string) (*view.VCenterInventoryView, error) {
	var resp view.VCenterInventoryView
	if err := cli.Get(ctx, "v1/vcenters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenter Pagination
func (cli *ZSClient) PageVCenter(ctx context.Context, params *param.QueryParam) ([]view.VCenterInventoryView, int, error) {
	var vCenters []view.VCenterInventoryView
	total, err := cli.Page(ctx, "v1/vcenters", params, &vCenters)
	return vCenters, total, err
}
// UpdateVCenter updates VCenter
func (cli *ZSClient) UpdateVCenter(ctx context.Context, uuid string, params param.UpdateVCenterParam) (*view.VCenterInventoryView, error) {
	resp := view.VCenterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vcenters", uuid, "", map[string]interface{}{
		"updateVCenter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVCenter deletes VCenter
func (cli *ZSClient) DeleteVCenter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vcenters", uuid, string(deleteMode))
}
