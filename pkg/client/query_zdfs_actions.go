// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryZdfs queries Zdfs list
func (cli *ZSClient) QueryZdfs(params *param.QueryParam) ([]view.ZdfsInventoryView, error) {
	var resp []view.ZdfsInventoryView
	return resp, cli.List("v1/zdfs", params, &resp)
}
