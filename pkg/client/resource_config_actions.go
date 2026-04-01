// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
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

// PageResourceConfig Pagination
func (cli *ZSClient) PageResourceConfig(params *param.QueryParam) ([]view.ResourceConfigInventoryView, int, error) {
	var resourceConfigs []view.ResourceConfigInventoryView
	total, err := cli.Page("v1/resource-configurations", params, &resourceConfigs)
	return resourceConfigs, total, err
}
// UpdateResourceConfig updates ResourceConfig
func (cli *ZSClient) UpdateResourceConfig(category string, name string, resourceUuid string, params param.UpdateResourceConfigParam) (*view.ResourceConfigInventoryView, error) {
	resp := view.ResourceConfigInventoryView{}
	err := cli.PutWithSpec("v1/resource-configurations", category, fmt.Sprintf("%s/%s/actions", name, resourceUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteResourceConfig deletes ResourceConfig
func (cli *ZSClient) DeleteResourceConfig(category string, name string, resourceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/resource-configurations", category, fmt.Sprintf("%s/%s", name, resourceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
