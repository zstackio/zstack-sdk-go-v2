// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModel queries Model list
func (cli *ZSClient) QueryModel(params *param.QueryParam) ([]view.ModelInventoryView, error) {
	var resp []view.ModelInventoryView
	return resp, cli.List("v1/ai/models", params, &resp)
}

func (cli *ZSClient) GetModel(uuid string) (*view.ModelInventoryView, error) {
	var resp view.ModelInventoryView
	if err := cli.Get("v1/ai/models", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteModel deletes Model
func (cli *ZSClient) DeleteModel(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/ai/models", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateModel updates Model
func (cli *ZSClient) UpdateModel(uuid string, params param.UpdateModelParam) (*view.ModelInventoryView, error) {
	var resp view.UpdateModelEventView
	err := cli.PutWithSpec("v1/ai/models", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddModel adds Model
func (cli *ZSClient) AddModel(params param.AddModelParam) (*view.ModelInventoryView, error) {
	var resp view.AddModelEventView
	if err := cli.Post("v1/ai/models", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddModelAsync Async
func (cli *ZSClient) AddModelAsync(params param.AddModelParam) (string, error) {

	resource := "v1/ai/models"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
