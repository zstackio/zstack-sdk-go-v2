// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedBlockGroupPrimaryStorage queries SharedBlockGroupPrimaryStorage list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorage(params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/sharedblockgroup", params, &resp)
}

// PageSharedBlockGroupPrimaryStorage Pagination
func (cli *ZSClient) PageSharedBlockGroupPrimaryStorage(params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageInventoryView, int, error) {
	var sharedBlockGroupPrimaryStorages []view.SharedBlockGroupPrimaryStorageInventoryView
	total, err := cli.Page("v1/primary-storage/sharedblockgroup", params, &sharedBlockGroupPrimaryStorages)
	return sharedBlockGroupPrimaryStorages, total, err
}
// AddSharedBlockGroupPrimaryStorage adds SharedBlockGroupPrimaryStorage
func (cli *ZSClient) AddSharedBlockGroupPrimaryStorage(params param.AddSharedBlockGroupPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/sharedblockgroup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
