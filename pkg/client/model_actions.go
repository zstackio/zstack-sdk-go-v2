// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModel queries Model list
func (cli *ZSClient) QueryModel(ctx context.Context, params *param.QueryParam) ([]view.ModelInventoryView, error) {
	var resp []view.ModelInventoryView
	return resp, cli.List(ctx, "v1/ai/models", params, &resp)
}

func (cli *ZSClient) GetModel(ctx context.Context, uuid string) (*view.ModelInventoryView, error) {
	var resp view.ModelInventoryView
	if err := cli.Get(ctx, "v1/ai/models", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModel Pagination
func (cli *ZSClient) PageModel(ctx context.Context, params *param.QueryParam) ([]view.ModelInventoryView, int, error) {
	var models []view.ModelInventoryView
	total, err := cli.Page(ctx, "v1/ai/models", params, &models)
	return models, total, err
}
// DeleteModel deletes Model
func (cli *ZSClient) DeleteModel(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/models", uuid, string(deleteMode))
}
// UpdateModel updates Model
func (cli *ZSClient) UpdateModel(ctx context.Context, uuid string, params param.UpdateModelParam) (*view.ModelInventoryView, error) {
	resp := view.ModelInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/models", uuid, "", map[string]interface{}{
		"updateModel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddModel adds Model
func (cli *ZSClient) AddModel(ctx context.Context, params param.AddModelParam) (*view.ModelInventoryView, error) {
	resp := view.ModelInventoryView{}
	if err := cli.Post(ctx, "v1/ai/models", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddModelAsync Async
func (cli *ZSClient) AddModelAsync(ctx context.Context, params param.AddModelParam) (string, error) {

	resource := "v1/ai/models"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
