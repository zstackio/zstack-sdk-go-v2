// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateXskyBlockVolume updates XskyBlockVolume
func (cli *ZSClient) UpdateXskyBlockVolume(uuid string, params param.UpdateXskyBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	var resp view.UpdateBlockVolumeEventView
	if err := cli.Put("v1/xsky/block-volumes", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryXskyBlockVolume queries XskyBlockVolume list
func (cli *ZSClient) QueryXskyBlockVolume(params *param.QueryParam) ([]view.XskyBlockVolumeInventoryView, error) {
	var resp []view.XskyBlockVolumeInventoryView
	return resp, cli.List("v1/xksy/block-volumes", params, &resp)
}

func (cli *ZSClient) GetXskyBlockVolume(uuid string) (*view.XskyBlockVolumeInventoryView, error) {
	var resp view.XskyBlockVolumeInventoryView
	if err := cli.Get("v1/xksy/block-volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
