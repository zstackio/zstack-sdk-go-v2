// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteResourceStack deletes ResourceStack
func (cli *ZSClient) DeleteResourceStack(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/cloudformation/stack", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryResourceStack queries ResourceStack list
func (cli *ZSClient) QueryResourceStack(params *param.QueryParam) ([]view.ResourceStackInventoryView, error) {
	var resp []view.ResourceStackInventoryView
	return resp, cli.List("v1/cloudformation/stack", params, &resp)
}

func (cli *ZSClient) GetResourceStack(uuid string) (*view.ResourceStackInventoryView, error) {
	var resp view.ResourceStackInventoryView
	if err := cli.Get("v1/cloudformation/stack", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	err := cli.PutWithSpec("v1/cloudformation/stack", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
