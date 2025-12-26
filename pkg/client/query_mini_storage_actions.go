// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMiniStorage queries MiniStorage list
func (cli *ZSClient) QueryMiniStorage(params *param.QueryParam) ([]view.MiniStorageInventoryView, error) {
	var resp []view.MiniStorageInventoryView
	return resp, cli.List("v1/primary-storage/mini", params, &resp)
}
