// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVolume queries Volume list
func (cli *ZSClient) QueryVolume(params *param.QueryParam) ([]view.VolumeInventoryView, error) {
	var resp []view.VolumeInventoryView
	return resp, cli.List("v1/volumes", params, &resp)
}
