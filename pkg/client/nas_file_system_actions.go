// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNasFileSystem queries NasFileSystem list
func (cli *ZSClient) QueryNasFileSystem(params *param.QueryParam) ([]view.NasFileSystemInventoryView, error) {
	var resp []view.NasFileSystemInventoryView
	return resp, cli.List("v1/primary-storage/nas", params, &resp)
}

func (cli *ZSClient) GetNasFileSystem(uuid string) (*view.NasFileSystemInventoryView, error) {
	var resp view.NasFileSystemInventoryView
	if err := cli.Get("v1/primary-storage/nas", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateNasFileSystem updates NasFileSystem
func (cli *ZSClient) UpdateNasFileSystem(uuid string, params param.UpdateNasFileSystemParam) (*view.NasFileSystemInventoryView, error) {
	var resp view.UpdateNasFileSystemEventView
	err := cli.PutWithSpec("v1/primary-storage/nas", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteNasFileSystem deletes NasFileSystem
func (cli *ZSClient) DeleteNasFileSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/primary-storage/nas", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
