// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateXskyBlockVolume updates XskyBlockVolume
func (cli *ZSClient) UpdateXskyBlockVolume(uuid string, params param.UpdateXskyBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	var resp view.UpdateBlockVolumeEventView
	if err := cli.Put("v1/xsky/block-volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryXskyBlockVolume queries XskyBlockVolume list
func (cli *ZSClient) QueryXskyBlockVolume(params *param.QueryParam) ([]view.XskyBlockVolumeInventoryView, error) {
	var resp []view.XskyBlockVolumeInventoryView
	return resp, cli.List("v1/xksy/block-volumes", params, &resp)
}
