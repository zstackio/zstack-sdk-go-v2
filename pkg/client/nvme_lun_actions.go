// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNvmeLun queries NvmeLun list
func (cli *ZSClient) QueryNvmeLun(params *param.QueryParam) ([]view.NvmeLunInventoryView, error) {
	var resp []view.NvmeLunInventoryView
	return resp, cli.List("v1/storage-devices/nvme/luns", params, &resp)
}

func (cli *ZSClient) GetNvmeLun(uuid string) (*view.NvmeLunInventoryView, error) {
	var resp view.NvmeLunInventoryView
	if err := cli.Get("v1/storage-devices/nvme/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
