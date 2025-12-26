// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVCenter queries VCenter list
func (cli *ZSClient) QueryVCenter(params *param.QueryParam) ([]view.VCenterInventoryView, error) {
	var resp []view.VCenterInventoryView
	return resp, cli.List("v1/vcenters", params, &resp)
}
