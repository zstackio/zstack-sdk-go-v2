// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHost queries Host list
func (cli *ZSClient) QueryHost(params *param.QueryParam) ([]view.HostInventoryView, error) {
	var resp []view.HostInventoryView
	return resp, cli.List("v1/hosts", params, &resp)
}
