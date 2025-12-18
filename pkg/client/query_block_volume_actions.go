// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBlockVolume queries BlockVolume list
func (cli *ZSClient) QueryBlockVolume(params param.QueryParam) ([]view.BlockVolumeInventoryView, error) {
	var resp []view.BlockVolumeInventoryView
	return resp, cli.List("v1/block-volumes", &params, &resp)
}
