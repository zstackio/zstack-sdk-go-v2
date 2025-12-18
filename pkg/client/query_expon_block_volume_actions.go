// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryExponBlockVolume queries ExponBlockVolume list
func (cli *ZSClient) QueryExponBlockVolume(params param.QueryParam) ([]view.ExponBlockVolumeInventoryView, error) {
	var resp []view.ExponBlockVolumeInventoryView
	return resp, cli.List("v1/expon/block-volumes", &params, &resp)
}
