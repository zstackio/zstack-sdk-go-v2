// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModel queries Model list
func (cli *ZSClient) QueryModel(params *param.QueryParam) ([]view.ModelInventoryView, error) {
	var resp []view.ModelInventoryView
	return resp, cli.List("v1/ai/models", params, &resp)
}
// DeleteModel deletes Model
func (cli *ZSClient) DeleteModel(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models/{uuid}", uuid, string(deleteMode))
}
// UpdateModel updates Model
func (cli *ZSClient) UpdateModel(uuid string, params param.UpdateModelParam) (*view.ModelInventoryView, error) {
	var resp view.UpdateModelEventView
	if err := cli.Put("v1/ai/models/{uuid}", uuid, params, &resp); err != nil {
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
