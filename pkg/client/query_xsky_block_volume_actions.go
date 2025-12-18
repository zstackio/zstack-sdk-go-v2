// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryXskyBlockVolume queries XskyBlockVolume list
func (cli *ZSClient) QueryXskyBlockVolume(params param.QueryParam) ([]view.XskyBlockVolumeInventoryView, error) {
	var resp []view.XskyBlockVolumeInventoryView
	return resp, cli.List("v1/xksy/block-volumes", &params, &resp)
}
