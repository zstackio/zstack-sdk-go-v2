// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNvmeServer queries NvmeServer list
func (cli *ZSClient) QueryNvmeServer(params *param.QueryParam) ([]view.NvmeServerInventoryView, error) {
	var resp []view.NvmeServerInventoryView
	return resp, cli.List("v1/storage-devices/nvme/servers", params, &resp)
}

func (cli *ZSClient) GetNvmeServer(uuid string) (*view.NvmeServerInventoryView, error) {
	var resp view.NvmeServerInventoryView
	if err := cli.Get("v1/storage-devices/nvme/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteNvmeServer deletes NvmeServer
func (cli *ZSClient) DeleteNvmeServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/storage-devices/nvme/servers", uuid, string(deleteMode))
}
// AddNvmeServer adds NvmeServer
func (cli *ZSClient) AddNvmeServer(params param.AddNvmeServerParam) (*view.NvmeServerInventoryView, error) {
	var resp view.AddNvmeServerEventView
	if err := cli.Post("v1/storage-devices/nvme/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
