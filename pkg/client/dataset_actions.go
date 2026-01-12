// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryDataset queries Dataset list
func (cli *ZSClient) QueryDataset(params *param.QueryParam) ([]view.DatasetInventoryView, error) {
	var resp []view.DatasetInventoryView
	return resp, cli.List("v1/ai/datasets", params, &resp)
}

func (cli *ZSClient) GetDataset(uuid string) (*view.DatasetInventoryView, error) {
	var resp view.DatasetInventoryView
	if err := cli.Get("v1/ai/datasets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateDataset creates Dataset
func (cli *ZSClient) CreateDataset(params param.CreateDatasetParam) (*view.DatasetInventoryView, error) {
	var resp view.CreateDatasetEventView
	if err := cli.Post("v1/ai/datasets", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateDatasetAsync Async
func (cli *ZSClient) CreateDatasetAsync(params param.CreateDatasetParam) (string, error) {

	resource := "v1/ai/datasets"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// DeleteDataset deletes Dataset
func (cli *ZSClient) DeleteDataset(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/ai/datasets", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateDataset updates Dataset
func (cli *ZSClient) UpdateDataset(uuid string, params param.UpdateDatasetParam) (*view.DatasetInventoryView, error) {
	var resp view.UpdateDatasetEventView
	err := cli.PutWithSpec("v1/ai/datasets", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
