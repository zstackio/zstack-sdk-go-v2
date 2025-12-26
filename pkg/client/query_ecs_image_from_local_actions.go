// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEcsImageFromLocal queries EcsImageFromLocal list
func (cli *ZSClient) QueryEcsImageFromLocal(params *param.QueryParam) ([]view.EcsImageInventoryView, error) {
	var resp []view.EcsImageInventoryView
	return resp, cli.List("v1/hybrid/aliyun/image", params, &resp)
}
