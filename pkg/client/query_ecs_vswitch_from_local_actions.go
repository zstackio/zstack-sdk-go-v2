// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEcsVSwitchFromLocal queries EcsVSwitchFromLocal list
func (cli *ZSClient) QueryEcsVSwitchFromLocal(params *param.QueryParam) ([]view.EcsVSwitchInventoryView, error) {
	var resp []view.EcsVSwitchInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vswitch", params, &resp)
}
