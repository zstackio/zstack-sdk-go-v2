// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHostPhysicalMemory queries HostPhysicalMemory list
func (cli *ZSClient) QueryHostPhysicalMemory(params *param.QueryParam) ([]view.HostPhysicalMemoryInventoryView, error) {
	var resp []view.HostPhysicalMemoryInventoryView
	return resp, cli.List("v1/hosts/physicalmemory", params, &resp)
}
