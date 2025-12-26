// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBareMetal2Instance queries BareMetal2Instance list
func (cli *ZSClient) QueryBareMetal2Instance(params *param.QueryParam) ([]view.BareMetal2InstanceInventoryView, error) {
	var resp []view.BareMetal2InstanceInventoryView
	return resp, cli.List("v1/baremetal2/bm-instances", params, &resp)
}
