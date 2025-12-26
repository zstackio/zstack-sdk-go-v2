// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBlockPrimaryStorage queries BlockPrimaryStorage list
func (cli *ZSClient) QueryBlockPrimaryStorage(params *param.QueryParam) ([]view.BlockPrimaryStorageInventoryView, error) {
	var resp []view.BlockPrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/block", params, &resp)
}
