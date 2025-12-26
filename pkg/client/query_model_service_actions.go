// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryModelService queries ModelService list
func (cli *ZSClient) QueryModelService(params *param.QueryParam) ([]view.ModelServiceInventoryView, error) {
	var resp []view.ModelServiceInventoryView
	return resp, cli.List("v1/ai/model-services", params, &resp)
}
