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
func (cli *ZSClient) UpdateResourceConfig(uuid string, params param.UpdateResourceConfigParam) (*view.ResourceConfigInventoryView, error) {
	var resp view.UpdateResourceConfigEventView
	if err := cli.Put("v1/resource-configurations/{category}/{name}/{resourceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteResourceConfig deletes ResourceConfig
func (cli *ZSClient) DeleteResourceConfig(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/resource-configurations/{category}/{name}/{resourceUuid}", uuid, string(deleteMode))
}
