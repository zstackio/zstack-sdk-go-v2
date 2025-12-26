// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEcsVpcFromLocal queries EcsVpcFromLocal list
func (cli *ZSClient) QueryEcsVpcFromLocal(params *param.QueryParam) ([]view.EcsVpcInventoryView, error) {
	var resp []view.EcsVpcInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vpc", params, &resp)
}
