// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryLocalStorageResourceRef queries LocalStorageResourceRef list
func (cli *ZSClient) QueryLocalStorageResourceRef(params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, error) {
	var resp []view.LocalStorageResourceRefInventoryView
	return resp, cli.List("v1/primary-storage/local-storage/resource-refs", params, &resp)
}
