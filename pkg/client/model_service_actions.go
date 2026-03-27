// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CloneModelService operates on ModelService
func (cli *ZSClient) CloneModelService(ctx context.Context, params param.CloneModelServiceParam) (*view.ModelServiceInventoryView, error) {
	resp := view.ModelServiceInventoryView{}
	if err := cli.Post(ctx, "v1/ai/model-services/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateModelService updates ModelService
func (cli *ZSClient) UpdateModelService(ctx context.Context, uuid string, params param.UpdateModelServiceParam) (*view.ModelServiceInventoryView, error) {
	resp := view.ModelServiceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/model-services", uuid, "", map[string]interface{}{
		"updateModelService": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteModelService deletes ModelService
func (cli *ZSClient) DeleteModelService(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/model-services", uuid, string(deleteMode))
}
// AddModelService adds ModelService
func (cli *ZSClient) AddModelService(ctx context.Context, params param.AddModelServiceParam) (*view.ModelServiceInventoryView, error) {
	resp := view.ModelServiceInventoryView{}
	if err := cli.Post(ctx, "v1/ai/model-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddModelServiceAsync Async
func (cli *ZSClient) AddModelServiceAsync(ctx context.Context, params param.AddModelServiceParam) (string, error) {

	resource := "v1/ai/model-services"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// QueryModelService queries ModelService list
func (cli *ZSClient) QueryModelService(ctx context.Context, params *param.QueryParam) ([]view.ModelServiceInventoryView, error) {
	var resp []view.ModelServiceInventoryView
	return resp, cli.List(ctx, "v1/ai/model-services", params, &resp)
}

func (cli *ZSClient) GetModelService(ctx context.Context, uuid string) (*view.ModelServiceInventoryView, error) {
	var resp view.ModelServiceInventoryView
	if err := cli.Get(ctx, "v1/ai/model-services", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModelService Pagination
func (cli *ZSClient) PageModelService(ctx context.Context, params *param.QueryParam) ([]view.ModelServiceInventoryView, int, error) {
	var modelServices []view.ModelServiceInventoryView
	total, err := cli.Page(ctx, "v1/ai/model-services", params, &modelServices)
	return modelServices, total, err
}
