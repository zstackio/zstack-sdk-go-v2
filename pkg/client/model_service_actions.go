// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CloneModelService operates on ModelService
func (cli *ZSClient) CloneModelService(uuid string, params param.CloneModelServiceParam) (*view.ModelServiceInventoryView, error) {
	var resp view.CloneModelServiceEventView
	err := cli.PostWithSpec("v1/ai/model-services", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateModelService updates ModelService
func (cli *ZSClient) UpdateModelService(uuid string, params param.UpdateModelServiceParam) (*view.ModelServiceInventoryView, error) {
	var resp view.UpdateModelServiceEventView
	err := cli.PutWithSpec("v1/ai/model-services", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteModelService deletes ModelService
func (cli *ZSClient) DeleteModelService(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/ai/model-services", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// AddModelService adds ModelService
func (cli *ZSClient) AddModelService(params param.AddModelServiceParam) (*view.ModelServiceInventoryView, error) {
	var resp view.AddModelServiceEventView
	if err := cli.Post("v1/ai/model-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// AddModelServiceAsync Async
func (cli *ZSClient) AddModelServiceAsync(params param.AddModelServiceParam) (string, error) {

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
func (cli *ZSClient) QueryModelService(params *param.QueryParam) ([]view.ModelServiceInventoryView, error) {
	var resp []view.ModelServiceInventoryView
	return resp, cli.List("v1/ai/model-services", params, &resp)
}

func (cli *ZSClient) GetModelService(uuid string) (*view.ModelServiceInventoryView, error) {
	var resp view.ModelServiceInventoryView
	if err := cli.Get("v1/ai/model-services", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
