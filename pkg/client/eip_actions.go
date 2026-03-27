// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateEip creates Eip
func (cli *ZSClient) CreateEip(ctx context.Context, params param.CreateEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.Post(ctx, "v1/eips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AttachEip operates on Eip
func (cli *ZSClient) AttachEip(ctx context.Context, params param.AttachEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.Post(ctx, "v1/eips/{eipUuid}/vm-instances/nics/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateEip updates Eip
func (cli *ZSClient) UpdateEip(ctx context.Context, uuid string, params param.UpdateEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/eips", uuid, "", map[string]interface{}{
		"updateEip": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryEip queries Eip list
func (cli *ZSClient) QueryEip(ctx context.Context, params *param.QueryParam) ([]view.EipInventoryView, error) {
	var resp []view.EipInventoryView
	return resp, cli.List(ctx, "v1/eips", params, &resp)
}

func (cli *ZSClient) GetEip(ctx context.Context, uuid string) (*view.EipInventoryView, error) {
	var resp view.EipInventoryView
	if err := cli.Get(ctx, "v1/eips", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEip Pagination
func (cli *ZSClient) PageEip(ctx context.Context, params *param.QueryParam) ([]view.EipInventoryView, int, error) {
	var eips []view.EipInventoryView
	total, err := cli.Page(ctx, "v1/eips", params, &eips)
	return eips, total, err
}
// DeleteEip deletes Eip
func (cli *ZSClient) DeleteEip(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/eips", uuid, string(deleteMode))
}
// DetachEip operates on Eip
func (cli *ZSClient) DetachEip(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/eips", uuid, string(deleteMode))
}
