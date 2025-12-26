// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryUserProxyConfig queries UserProxyConfig list
func (cli *ZSClient) QueryUserProxyConfig(params *param.QueryParam) ([]view.UserProxyConfigInventoryView, error) {
	var resp []view.UserProxyConfigInventoryView
	return resp, cli.List("v1/user-proxy-configs", params, &resp)
}
