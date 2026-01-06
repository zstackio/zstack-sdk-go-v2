// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryDataset queries Dataset list
func (cli *ZSClient) QueryDataset(params *param.QueryParam) ([]view.DatasetInventoryView, error) {
	var resp []view.DatasetInventoryView
	return resp, cli.List("v1/ai/datasets", params, &resp)
}
// CreateDataset creates Dataset
func (cli *ZSClient) CreateDataset(params param.CreateDatasetParam) (*view.DatasetInventoryView, error) {
	var resp view.CreateDatasetEventView
	if err := cli.Post("v1/ai/datasets", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteDataset deletes Dataset
func (cli *ZSClient) DeleteDataset(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/datasets/{uuid}", uuid, string(deleteMode))
}
// UpdateDataset updates Dataset
func (cli *ZSClient) UpdateDataset(uuid string, params param.UpdateDatasetParam) (*view.DatasetInventoryView, error) {
	var resp view.UpdateDatasetEventView
	if err := cli.Put("v1/ai/datasets/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
