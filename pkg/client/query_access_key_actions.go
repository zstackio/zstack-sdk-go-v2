// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAccessKey queries AccessKey list
func (cli *ZSClient) QueryAccessKey(params *param.QueryParam) ([]view.AccessKeyInventoryView, error) {
	var resp []view.AccessKeyInventoryView
	return resp, cli.List("v1/accesskeys", params, &resp)
}
