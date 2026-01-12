// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBlockPrimaryStorage updates BlockPrimaryStorage
func (cli *ZSClient) UpdateBlockPrimaryStorage(uuid string, params param.UpdateBlockPrimaryStorageParam) (*view.BlockPrimaryStorageInventoryView, error) {
	var resp view.UpdateBlockPrimaryStorageEventView
	if err := cli.Put("v1/primary-storage/block", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddBlockPrimaryStorage adds BlockPrimaryStorage
func (cli *ZSClient) AddBlockPrimaryStorage(params param.AddBlockPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/block", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBlockPrimaryStorage queries BlockPrimaryStorage list
func (cli *ZSClient) QueryBlockPrimaryStorage(params *param.QueryParam) ([]view.BlockPrimaryStorageInventoryView, error) {
	var resp []view.BlockPrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/block", params, &resp)
}

func (cli *ZSClient) GetBlockPrimaryStorage(uuid string) (*view.BlockPrimaryStorageInventoryView, error) {
	var resp view.BlockPrimaryStorageInventoryView
	if err := cli.Get("v1/primary-storage/block", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
