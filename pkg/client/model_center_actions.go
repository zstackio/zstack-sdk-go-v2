// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteModelCenter deletes ModelCenter
func (cli *ZSClient) DeleteModelCenter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/model-centers", uuid, string(deleteMode))
}
// QueryModelCenter queries ModelCenter list
func (cli *ZSClient) QueryModelCenter(ctx context.Context, params *param.QueryParam) ([]view.ModelCenterInventoryView, error) {
	var resp []view.ModelCenterInventoryView
	return resp, cli.List(ctx, "v1/ai/model-centers", params, &resp)
}

func (cli *ZSClient) GetModelCenter(ctx context.Context, uuid string) (*view.ModelCenterInventoryView, error) {
	var resp view.ModelCenterInventoryView
	if err := cli.Get(ctx, "v1/ai/model-centers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModelCenter Pagination
func (cli *ZSClient) PageModelCenter(ctx context.Context, params *param.QueryParam) ([]view.ModelCenterInventoryView, int, error) {
	var modelCenters []view.ModelCenterInventoryView
	total, err := cli.Page(ctx, "v1/ai/model-centers", params, &modelCenters)
	return modelCenters, total, err
}
// AddModelCenter adds ModelCenter
func (cli *ZSClient) AddModelCenter(ctx context.Context, params param.AddModelCenterParam) (*view.ModelCenterInventoryView, error) {
	resp := view.ModelCenterInventoryView{}
	if err := cli.Post(ctx, "v1/ai/model-centers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateModelCenter updates ModelCenter
func (cli *ZSClient) UpdateModelCenter(ctx context.Context, uuid string, params param.UpdateModelCenterParam) (*view.ModelCenterInventoryView, error) {
	resp := view.ModelCenterInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/model-centers", uuid, "", map[string]interface{}{
		"updateModelCenter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
