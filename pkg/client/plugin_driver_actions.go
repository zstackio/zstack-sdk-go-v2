// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryPluginDrivers queries PluginDrivers list
func (cli *ZSClient) QueryPluginDrivers(params *param.QueryParam) ([]view.PluginDriverInventoryView, error) {
	var resp []view.PluginDriverInventoryView
	return resp, cli.List("v1/external/plugins", params, &resp)
}

func (cli *ZSClient) GetPluginDrivers(uuid string) (*view.PluginDriverInventoryView, error) {
	var resp view.PluginDriverInventoryView
	if err := cli.Get("v1/external/plugins", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePluginDrivers Pagination
func (cli *ZSClient) PagePluginDrivers(params *param.QueryParam) ([]view.PluginDriverInventoryView, int, error) {
	var pluginDrivers []view.PluginDriverInventoryView
	total, err := cli.Page("v1/external/plugins", params, &pluginDrivers)
	return pluginDrivers, total, err
}

// RefreshPluginDrivers operates on PluginDrivers
func (cli *ZSClient) RefreshPluginDrivers(params param.RefreshPluginDriversParam) (*view.RefreshPluginDriversEventView, error) {
	resp := view.RefreshPluginDriversEventView{}
	if err := cli.PutWithRespKey("v1/external/plugins", "", "", map[string]interface{}{
		"refreshPluginDrivers": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePluginDrivers deletes PluginDrivers
func (cli *ZSClient) DeletePluginDrivers(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/external/plugins", uuid, string(deleteMode))
}
