// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBlockVolume updates BlockVolume
func (cli *ZSClient) UpdateBlockVolume(uuid string, params param.UpdateBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	var resp view.UpdateBlockVolumeEventView
	if err := cli.Put("v1/block-volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateBlockVolume creates BlockVolume
func (cli *ZSClient) CreateBlockVolume(params param.CreateBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	var resp view.CreateBlockVolumeEventView
	if err := cli.Post("v1/block-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBlockVolume queries BlockVolume list
func (cli *ZSClient) QueryBlockVolume(params *param.QueryParam) ([]view.BlockVolumeInventoryView, error) {
	var resp []view.BlockVolumeInventoryView
	return resp, cli.List("v1/block-volumes", params, &resp)
}
