// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryModel queries Model list
func (cli *ZSClient) QueryModel(params *param.QueryParam) ([]view.ModelInventoryView, error) {
	var resp []view.ModelInventoryView
	return resp, cli.List("v1/ai/models", params, &resp)
}
