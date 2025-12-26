// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySharedBlockGroupPrimaryStorage queries SharedBlockGroupPrimaryStorage list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorage(params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/sharedblockgroup", params, &resp)
}
