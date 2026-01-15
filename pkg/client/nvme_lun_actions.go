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

// PageNvmeLun Pagination
func (cli *ZSClient) PageNvmeLun(params *param.QueryParam) ([]view.NvmeLunInventoryView, int, error) {
	var nvmeLuns []view.NvmeLunInventoryView
	total, err := cli.Page("v1/storage-devices/nvme/luns", params, &nvmeLuns)
	return nvmeLuns, total, err
}
