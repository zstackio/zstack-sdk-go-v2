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
	return cli.Delete("v1/cloudformation/stack", uuid, string(deleteMode))
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

// PageResourceStack Pagination
func (cli *ZSClient) PageResourceStack(params *param.QueryParam) ([]view.ResourceStackInventoryView, int, error) {
	var resourceStacks []view.ResourceStackInventoryView
	total, err := cli.Page("v1/cloudformation/stack", params, &resourceStacks)
	return resourceStacks, total, err
}
// CreateResourceStack creates ResourceStack
func (cli *ZSClient) CreateResourceStack(params param.CreateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.Post("v1/cloudformation/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateResourceStack updates ResourceStack
func (cli *ZSClient) UpdateResourceStack(uuid string, params param.UpdateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.PutWithRespKey("v1/cloudformation/stack", uuid, "", map[string]interface{}{
		"updateResourceStack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
