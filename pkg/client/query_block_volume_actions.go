// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBlockVolume queries BlockVolume list
func (cli *ZSClient) QueryBlockVolume(params *param.QueryParam) ([]view.BlockVolumeInventoryView, error) {
	var resp []view.BlockVolumeInventoryView
	return resp, cli.List("v1/block-volumes", params, &resp)
}
