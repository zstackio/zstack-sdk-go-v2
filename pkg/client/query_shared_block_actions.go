// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySharedBlock queries SharedBlock list
func (cli *ZSClient) QuerySharedBlock(params *param.QueryParam) ([]view.SharedBlockInventoryView, error) {
	var resp []view.SharedBlockInventoryView
	return resp, cli.List("v1/sharedblock-group/sharedblocks", params, &resp)
}
