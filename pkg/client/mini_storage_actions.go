// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMiniStorage queries MiniStorage list
func (cli *ZSClient) QueryMiniStorage(params *param.QueryParam) ([]view.MiniStorageInventoryView, error) {
	var resp []view.MiniStorageInventoryView
	return resp, cli.List("v1/primary-storage/mini", params, &resp)
}

// PageMiniStorage Pagination
func (cli *ZSClient) PageMiniStorage(params *param.QueryParam) ([]view.MiniStorageInventoryView, int, error) {
	var miniStorages []view.MiniStorageInventoryView
	total, err := cli.Page("v1/primary-storage/mini", params, &miniStorages)
	return miniStorages, total, err
}
// AddMiniStorage adds MiniStorage
func (cli *ZSClient) AddMiniStorage(params param.AddMiniStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/mini", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
