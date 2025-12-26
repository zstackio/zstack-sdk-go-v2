// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPluginDrivers queries PluginDrivers list
func (cli *ZSClient) QueryPluginDrivers(params *param.QueryParam) ([]view.PluginDriverInventoryView, error) {
	var resp []view.PluginDriverInventoryView
	return resp, cli.List("v1/external/plugins", params, &resp)
}
