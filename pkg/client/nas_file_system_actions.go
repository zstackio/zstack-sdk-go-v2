// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNasFileSystem queries NasFileSystem list
func (cli *ZSClient) QueryNasFileSystem(params *param.QueryParam) ([]view.NasFileSystemInventoryView, error) {
	var resp []view.NasFileSystemInventoryView
	return resp, cli.List("v1/primary-storage/nas", params, &resp)
}
// UpdateNasFileSystem updates NasFileSystem
func (cli *ZSClient) UpdateNasFileSystem(uuid string, params param.UpdateNasFileSystemParam) (*view.NasFileSystemInventoryView, error) {
	var resp view.UpdateNasFileSystemEventView
	if err := cli.Put("v1/primary-storage/nas/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteNasFileSystem deletes NasFileSystem
func (cli *ZSClient) DeleteNasFileSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas/{uuid}", uuid, string(deleteMode))
}
