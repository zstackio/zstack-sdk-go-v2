// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteResourceStack deletes ResourceStack
func (cli *ZSClient) DeleteResourceStack(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/stack/{uuid}", uuid, string(deleteMode))
}
// QueryResourceStack queries ResourceStack list
func (cli *ZSClient) QueryResourceStack(params *param.QueryParam) ([]view.ResourceStackInventoryView, error) {
	var resp []view.ResourceStackInventoryView
	return resp, cli.List("v1/cloudformation/stack", params, &resp)
}
// CreateResourceStack creates ResourceStack
func (cli *ZSClient) CreateResourceStack(params param.CreateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	var resp view.CreateResourceStackEventView
	if err := cli.Post("v1/cloudformation/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateResourceStack updates ResourceStack
func (cli *ZSClient) UpdateResourceStack(uuid string, params param.UpdateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	var resp view.UpdateResourceStackEventView
	if err := cli.Put("v1/cloudformation/stack/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
