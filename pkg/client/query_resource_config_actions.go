// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryResourceConfig queries ResourceConfig list
func (cli *ZSClient) QueryResourceConfig(params *param.QueryParam) ([]view.ResourceConfigInventoryView, error) {
	var resp []view.ResourceConfigInventoryView
	return resp, cli.List("v1/resource-configurations", params, &resp)
}
