// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBaremetalInstance queries BaremetalInstance list
func (cli *ZSClient) QueryBaremetalInstance(params *param.QueryParam) ([]view.BaremetalInstanceInventoryView, error) {
	var resp []view.BaremetalInstanceInventoryView
	return resp, cli.List("v1/baremetal/instances", params, &resp)
}
