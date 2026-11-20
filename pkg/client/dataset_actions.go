// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"

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

// PageDataset Pagination
func (cli *ZSClient) PageDataset(params *param.QueryParam) ([]view.DatasetInventoryView, int, error) {
	var datasets []view.DatasetInventoryView
	total, err := cli.Page("v1/ai/datasets", params, &datasets)
	return datasets, total, err
}
// CreateDataset creates Dataset
func (cli *ZSClient) CreateDataset(params param.CreateDatasetParam) (*view.DatasetInventoryView, error) {
	resp := view.DatasetInventoryView{}
	if err := cli.Post("v1/ai/datasets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	return cli.Delete("v1/ai/datasets", uuid, string(deleteMode))
}
// UpdateDataset updates Dataset
func (cli *ZSClient) UpdateDataset(uuid string, params param.UpdateDatasetParam) (*view.DatasetInventoryView, error) {
	resp := view.DatasetInventoryView{}
	if err := cli.PutWithRespKey("v1/ai/datasets", uuid, "", map[string]interface{}{
		"updateDataset": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDatasets updates Datasets
func (cli *ZSClient) UpdateDatasets(params param.UpdateDatasetsParam) (*view.UpdateDatasetsEventView, error) {
	resp := view.UpdateDatasetsEventView{}
	if err := cli.PutWithRespKey("v1/ai/datasets", "", "", map[string]interface{}{
		"updateDatasets": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatasets sends one UUID to the collection endpoint.
func (cli *ZSClient) DeleteDatasets(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteDatasetsWithParams(deleteMode, param.DeleteDatasetsParamDetail{Uuids: []string{uuid}})
}

// DeleteDatasetsWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteDatasetsWithParams(deleteMode param.DeleteMode, params param.DeleteDatasetsParamDetail) error {
	query := url.Values{
		"uuids": params.Uuids,
	}
	return cli.deleteActionWithQuery("v1/ai/datasets", "", "", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "uuids")
}
