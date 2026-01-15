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
	resp := view.BlockPrimaryStorageInventoryView{}
	if err := cli.Put("v1/primary-storage/block", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddBlockPrimaryStorage adds BlockPrimaryStorage
func (cli *ZSClient) AddBlockPrimaryStorage(params param.AddBlockPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/block", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBlockPrimaryStorage queries BlockPrimaryStorage list
func (cli *ZSClient) QueryBlockPrimaryStorage(params *param.QueryParam) ([]view.BlockPrimaryStorageInventoryView, error) {
	var resp []view.BlockPrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/block", params, &resp)
}

// PageBlockPrimaryStorage Pagination
func (cli *ZSClient) PageBlockPrimaryStorage(params *param.QueryParam) ([]view.BlockPrimaryStorageInventoryView, int, error) {
	var blockPrimaryStorages []view.BlockPrimaryStorageInventoryView
	total, err := cli.Page("v1/primary-storage/block", params, &blockPrimaryStorages)
	return blockPrimaryStorages, total, err
}
