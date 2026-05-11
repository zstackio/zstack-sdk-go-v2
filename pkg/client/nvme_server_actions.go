// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNvmeServer queries NvmeServer list
func (cli *ZSClient) QueryNvmeServer(ctx context.Context, params *param.QueryParam) ([]view.NvmeServerInventoryView, error) {
	var resp []view.NvmeServerInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/nvme/servers", params, &resp)
}

func (cli *ZSClient) GetNvmeServer(ctx context.Context, uuid string) (*view.NvmeServerInventoryView, error) {
	var resp view.NvmeServerInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/nvme/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNvmeServer Pagination
func (cli *ZSClient) PageNvmeServer(ctx context.Context, params *param.QueryParam) ([]view.NvmeServerInventoryView, int, error) {
	var nvmeServers []view.NvmeServerInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/nvme/servers", params, &nvmeServers)
	return nvmeServers, total, err
}
// DeleteNvmeServer deletes NvmeServer
func (cli *ZSClient) DeleteNvmeServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/storage-devices/nvme/servers", uuid, string(deleteMode))
}
// AddNvmeServer adds NvmeServer
func (cli *ZSClient) AddNvmeServer(ctx context.Context, params param.AddNvmeServerParam) (*view.NvmeServerInventoryView, error) {
	resp := view.NvmeServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/storage-devices/nvme/servers", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
