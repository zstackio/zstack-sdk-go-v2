// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryDirectory queries Directory list
func (cli *ZSClient) QueryDirectory(params *param.QueryParam) ([]view.DirectoryInventoryView, error) {
	var resp []view.DirectoryInventoryView
	return resp, cli.List("v1/directories", params, &resp)
}
