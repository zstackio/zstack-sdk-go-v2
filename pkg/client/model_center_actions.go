// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteModelCenter deletes ModelCenter
func (cli *ZSClient) DeleteModelCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-centers/{uuid}", uuid, string(deleteMode))
}
// QueryModelCenter queries ModelCenter list
func (cli *ZSClient) QueryModelCenter(params *param.QueryParam) ([]view.ModelCenterInventoryView, error) {
	var resp []view.ModelCenterInventoryView
	return resp, cli.List("v1/ai/model-centers", params, &resp)
}
// AddModelCenter adds ModelCenter
func (cli *ZSClient) AddModelCenter(params param.AddModelCenterParam) (*view.ModelCenterInventoryView, error) {
	var resp view.AddModelCenterEventView
	if err := cli.Post("v1/ai/model-centers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateModelCenter updates ModelCenter
func (cli *ZSClient) UpdateModelCenter(uuid string, params param.UpdateModelCenterParam) (*view.ModelCenterInventoryView, error) {
	var resp view.UpdateModelCenterEventView
	if err := cli.Put("v1/ai/model-centers/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
