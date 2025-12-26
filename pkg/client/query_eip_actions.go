// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEip queries Eip list
func (cli *ZSClient) QueryEip(params *param.QueryParam) ([]view.EipInventoryView, error) {
	var resp []view.EipInventoryView
	return resp, cli.List("v1/eips", params, &resp)
}
