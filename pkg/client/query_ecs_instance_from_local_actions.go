// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEcsInstanceFromLocal queries EcsInstanceFromLocal list
func (cli *ZSClient) QueryEcsInstanceFromLocal(params *param.QueryParam) ([]view.EcsInstanceInventoryView, error) {
	var resp []view.EcsInstanceInventoryView
	return resp, cli.List("v1/hybrid/aliyun/ecs", params, &resp)
}
