// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryDataset queries Dataset list
func (cli *ZSClient) QueryDataset(ctx context.Context, params *param.QueryParam) ([]view.DatasetInventoryView, error) {
	var resp []view.DatasetInventoryView
	return resp, cli.List(ctx, "v1/ai/datasets", params, &resp)
}

func (cli *ZSClient) GetDataset(ctx context.Context, uuid string) (*view.DatasetInventoryView, error) {
	var resp view.DatasetInventoryView
	if err := cli.Get(ctx, "v1/ai/datasets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDataset Pagination
func (cli *ZSClient) PageDataset(ctx context.Context, params *param.QueryParam) ([]view.DatasetInventoryView, int, error) {
	var datasets []view.DatasetInventoryView
	total, err := cli.Page(ctx, "v1/ai/datasets", params, &datasets)
	return datasets, total, err
}
// CreateDataset creates Dataset
func (cli *ZSClient) CreateDataset(ctx context.Context, params param.CreateDatasetParam) (*view.DatasetInventoryView, error) {
	resp := view.DatasetInventoryView{}
	if err := cli.Post(ctx, "v1/ai/datasets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatasetAsync Async
func (cli *ZSClient) CreateDatasetAsync(ctx context.Context, params param.CreateDatasetParam) (string, error) {

	resource := "v1/ai/datasets"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// DeleteDataset deletes Dataset
func (cli *ZSClient) DeleteDataset(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ai/datasets", uuid, string(deleteMode))
}
// UpdateDataset updates Dataset
func (cli *ZSClient) UpdateDataset(ctx context.Context, uuid string, params param.UpdateDatasetParam) (*view.DatasetInventoryView, error) {
	resp := view.DatasetInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/ai/datasets", uuid, "", map[string]interface{}{
		"updateDataset": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
