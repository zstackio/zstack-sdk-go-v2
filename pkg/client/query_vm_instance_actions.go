// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmInstance queries VmInstance list
func (cli *ZSClient) QueryVmInstance(params *param.QueryParam) ([]view.VmInstanceInventoryView, error) {
	var resp []view.VmInstanceInventoryView
	return resp, cli.List("v1/vm-instances", params, &resp)
}
