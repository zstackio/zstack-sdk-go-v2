// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryDataset queries Dataset list
func (cli *ZSClient) QueryDataset(params *param.QueryParam) ([]view.DatasetInventoryView, error) {
	var resp []view.DatasetInventoryView
	return resp, cli.List("v1/ai/datasets", params, &resp)
}
