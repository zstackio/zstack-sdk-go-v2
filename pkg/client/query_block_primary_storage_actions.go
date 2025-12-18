// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBlockPrimaryStorage queries BlockPrimaryStorage list
func (cli *ZSClient) QueryBlockPrimaryStorage(params param.QueryParam) ([]view.BlockPrimaryStorageInventoryView, error) {
	var resp []view.BlockPrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/block", &params, &resp)
}
