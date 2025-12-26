// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryModelCenter queries ModelCenter list
func (cli *ZSClient) QueryModelCenter(params *param.QueryParam) ([]view.ModelCenterInventoryView, error) {
	var resp []view.ModelCenterInventoryView
	return resp, cli.List("v1/ai/model-centers", params, &resp)
}
