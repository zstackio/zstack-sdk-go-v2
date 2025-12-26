// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySlbVmInstance queries SlbVmInstance list
func (cli *ZSClient) QuerySlbVmInstance(params *param.QueryParam) ([]view.SlbVmInstanceInventoryView, error) {
	var resp []view.SlbVmInstanceInventoryView
	return resp, cli.List("v1/load-balancers/slb/instances", params, &resp)
}
