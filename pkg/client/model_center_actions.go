// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteModelCenter deletes ModelCenter
func (cli *ZSClient) DeleteModelCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/model-centers", uuid, string(deleteMode))
}
// QueryModelCenter queries ModelCenter list
func (cli *ZSClient) QueryModelCenter(params *param.QueryParam) ([]view.ModelCenterInventoryView, error) {
	var resp []view.ModelCenterInventoryView
	return resp, cli.List("v1/ai/model-centers", params, &resp)
}

func (cli *ZSClient) GetModelCenter(uuid string) (*view.ModelCenterInventoryView, error) {
	var resp view.ModelCenterInventoryView
	if err := cli.Get("v1/ai/model-centers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	if err := cli.Put("v1/ai/model-centers", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
