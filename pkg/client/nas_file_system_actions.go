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

// PageNasFileSystem Pagination
func (cli *ZSClient) PageNasFileSystem(params *param.QueryParam) ([]view.NasFileSystemInventoryView, int, error) {
	var nasFileSystems []view.NasFileSystemInventoryView
	total, err := cli.Page("v1/primary-storage/nas", params, &nasFileSystems)
	return nasFileSystems, total, err
}
// UpdateNasFileSystem updates NasFileSystem
func (cli *ZSClient) UpdateNasFileSystem(uuid string, params param.UpdateNasFileSystemParam) (*view.NasFileSystemInventoryView, error) {
	resp := view.NasFileSystemInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage/nas", uuid, "", map[string]interface{}{
		"updateNasFileSystem": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteNasFileSystem deletes NasFileSystem
func (cli *ZSClient) DeleteNasFileSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas", uuid, string(deleteMode))
}
