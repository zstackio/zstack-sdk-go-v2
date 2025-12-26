// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmPriorityConfig queries VmPriorityConfig list
func (cli *ZSClient) QueryVmPriorityConfig(params *param.QueryParam) ([]view.VmPriorityConfigInventoryView, error) {
	var resp []view.VmPriorityConfigInventoryView
	return resp, cli.List("v1/vm-priority-config", params, &resp)
}
