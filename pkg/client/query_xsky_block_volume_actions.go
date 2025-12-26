// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryXskyBlockVolume queries XskyBlockVolume list
func (cli *ZSClient) QueryXskyBlockVolume(params *param.QueryParam) ([]view.XskyBlockVolumeInventoryView, error) {
	var resp []view.XskyBlockVolumeInventoryView
	return resp, cli.List("v1/xksy/block-volumes", params, &resp)
}
