// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVtep queries Vtep list
func (cli *ZSClient) QueryVtep(params *param.QueryParam) ([]view.VtepInventoryView, error) {
	var resp []view.VtepInventoryView
	return resp, cli.List("v1/l2-networks/vteps", params, &resp)
}
