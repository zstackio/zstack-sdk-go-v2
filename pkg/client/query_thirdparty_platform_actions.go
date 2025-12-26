// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryThirdpartyPlatform queries ThirdpartyPlatform list
func (cli *ZSClient) QueryThirdpartyPlatform(params *param.QueryParam) ([]view.ThirdpartyPlatformInventoryView, error) {
	var resp []view.ThirdpartyPlatformInventoryView
	return resp, cli.List("v1/zwatch/third-party/platforms", params, &resp)
}
