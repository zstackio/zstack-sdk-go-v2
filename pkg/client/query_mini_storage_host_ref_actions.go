// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMiniStorageHostRef queries MiniStorageHostRef list
func (cli *ZSClient) QueryMiniStorageHostRef(params *param.QueryParam) ([]view.MiniStorageHostRefInventoryView, error) {
	var resp []view.MiniStorageHostRefInventoryView
	return resp, cli.List("v1/primary-storage/mini/host-refs", params, &resp)
}
