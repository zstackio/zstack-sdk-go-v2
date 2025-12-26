// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryGlobalConfig queries GlobalConfig list
func (cli *ZSClient) QueryGlobalConfig(params *param.QueryParam) ([]view.GlobalConfigInventoryView, error) {
	var resp []view.GlobalConfigInventoryView
	return resp, cli.List("v1/global-configurations", params, &resp)
}
