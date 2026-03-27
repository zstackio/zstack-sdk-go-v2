// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVpcSharedQos creates VpcSharedQos
func (cli *ZSClient) CreateVpcSharedQos(ctx context.Context, params param.CreateVpcSharedQosParam) (*view.VpcSharedQosInventoryView, error) {
	resp := view.VpcSharedQosInventoryView{}
	if err := cli.Post(ctx, "v1/vips/sharedqos", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVpcSharedQos queries VpcSharedQos list
func (cli *ZSClient) QueryVpcSharedQos(ctx context.Context, params *param.QueryParam) ([]view.VpcSharedQosInventoryView, error) {
	var resp []view.VpcSharedQosInventoryView
	return resp, cli.List(ctx, "v1/vips/sharedqos", params, &resp)
}

func (cli *ZSClient) GetVpcSharedQos(ctx context.Context, uuid string) (*view.VpcSharedQosInventoryView, error) {
	var resp view.VpcSharedQosInventoryView
	if err := cli.Get(ctx, "v1/vips/sharedqos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcSharedQos Pagination
func (cli *ZSClient) PageVpcSharedQos(ctx context.Context, params *param.QueryParam) ([]view.VpcSharedQosInventoryView, int, error) {
	var vpcSharedQos []view.VpcSharedQosInventoryView
	total, err := cli.Page(ctx, "v1/vips/sharedqos", params, &vpcSharedQos)
	return vpcSharedQos, total, err
}
// UpdateVpcSharedQos updates VpcSharedQos
func (cli *ZSClient) UpdateVpcSharedQos(ctx context.Context, uuid string, params param.UpdateVpcSharedQosParam) (*view.VpcSharedQosInventoryView, error) {
	resp := view.VpcSharedQosInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vips/sharedqos", uuid, "", map[string]interface{}{
		"updateVpcSharedQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVpcSharedQos deletes VpcSharedQos
func (cli *ZSClient) DeleteVpcSharedQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vips/sharedqos", uuid, string(deleteMode))
}
