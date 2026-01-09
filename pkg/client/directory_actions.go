// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateDirectory updates Directory
func (cli *ZSClient) UpdateDirectory(uuid string, params param.UpdateDirectoryParam) (*view.DirectoryInventoryView, error) {
	var resp view.UpdateDirectoryEventView
	if err := cli.Put("v1/update/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateDirectory creates Directory
func (cli *ZSClient) CreateDirectory(params param.CreateDirectoryParam) (*view.DirectoryInventoryView, error) {
	var resp view.CreateDirectoryEventView
	if err := cli.Post("v1/create/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryDirectory queries Directory list
func (cli *ZSClient) QueryDirectory(params *param.QueryParam) ([]view.DirectoryInventoryView, error) {
	var resp []view.DirectoryInventoryView
	return resp, cli.List("v1/directories", params, &resp)
}

func (cli *ZSClient) GetDirectory(uuid string) (*view.DirectoryInventoryView, error) {
	var resp view.DirectoryInventoryView
	if err := cli.Get("v1/directories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteDirectory deletes Directory
func (cli *ZSClient) DeleteDirectory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/directory", uuid, string(deleteMode))
}
