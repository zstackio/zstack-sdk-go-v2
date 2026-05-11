// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateDirectory updates Directory
func (cli *ZSClient) UpdateDirectory(ctx context.Context, params param.UpdateDirectoryParam) (*view.DirectoryInventoryView, error) {
	resp := view.DirectoryInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/update/directory", "", "", map[string]interface{}{
		"updateDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateDirectory creates Directory
func (cli *ZSClient) CreateDirectory(ctx context.Context, params param.CreateDirectoryParam) (*view.DirectoryInventoryView, error) {
	resp := view.DirectoryInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/create/directory", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryDirectory queries Directory list
func (cli *ZSClient) QueryDirectory(ctx context.Context, params *param.QueryParam) ([]view.DirectoryInventoryView, error) {
	var resp []view.DirectoryInventoryView
	return resp, cli.List(ctx, "v1/directories", params, &resp)
}

func (cli *ZSClient) GetDirectory(ctx context.Context, uuid string) (*view.DirectoryInventoryView, error) {
	var resp view.DirectoryInventoryView
	if err := cli.Get(ctx, "v1/directories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDirectory Pagination
func (cli *ZSClient) PageDirectory(ctx context.Context, params *param.QueryParam) ([]view.DirectoryInventoryView, int, error) {
	var directories []view.DirectoryInventoryView
	total, err := cli.Page(ctx, "v1/directories", params, &directories)
	return directories, total, err
}
// DeleteDirectory deletes Directory
func (cli *ZSClient) DeleteDirectory(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/delete/directory", uuid, string(deleteMode))
}
