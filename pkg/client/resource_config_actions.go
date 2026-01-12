// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryResourceConfig queries ResourceConfig list
func (cli *ZSClient) QueryResourceConfig(params *param.QueryParam) ([]view.ResourceConfigInventoryView, error) {
	var resp []view.ResourceConfigInventoryView
	return resp, cli.List("v1/resource-configurations", params, &resp)
}

func (cli *ZSClient) GetResourceConfig(uuid string) (*view.ResourceConfigInventoryView, error) {
	var resp view.ResourceConfigInventoryView
	if err := cli.Get("v1/resource-configurations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateResourceConfig updates ResourceConfig
func (cli *ZSClient) UpdateResourceConfig(category string, name string, resourceUuid string, params param.UpdateResourceConfigParam) (*view.ResourceConfigInventoryView, error) {
	var resp view.UpdateResourceConfigEventView
	err := cli.PutWithSpec("v1/resource-configurations", fmt.Sprintf(\"%s/%s/%s/actions\", category, name, resourceUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteResourceConfig deletes ResourceConfig
func (cli *ZSClient) DeleteResourceConfig(category string, name string, resourceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/resource-configurations", fmt.Sprintf(\"%s/%s/%s\", category, name, resourceUuid), string(deleteMode))
}
