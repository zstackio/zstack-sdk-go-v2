// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBlockVolume updates BlockVolume
func (cli *ZSClient) UpdateBlockVolume(uuid string, params param.UpdateBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	var resp view.UpdateBlockVolumeEventView
	if err := cli.Put("v1/block-volumes", uuid, params, &resp); err != nil {
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

func (cli *ZSClient) GetBlockVolume(uuid string) (*view.BlockVolumeInventoryView, error) {
	var resp view.BlockVolumeInventoryView
	if err := cli.Get("v1/block-volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
