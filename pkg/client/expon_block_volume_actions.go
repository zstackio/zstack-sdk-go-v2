// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryExponBlockVolume queries ExponBlockVolume list
func (cli *ZSClient) QueryExponBlockVolume(params *param.QueryParam) ([]view.ExponBlockVolumeInventoryView, error) {
	var resp []view.ExponBlockVolumeInventoryView
	return resp, cli.List("v1/expon/block-volumes", params, &resp)
}
