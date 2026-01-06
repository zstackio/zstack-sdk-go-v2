// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNvmeLun queries NvmeLun list
func (cli *ZSClient) QueryNvmeLun(params *param.QueryParam) ([]view.NvmeLunInventoryView, error) {
	var resp []view.NvmeLunInventoryView
	return resp, cli.List("v1/storage-devices/nvme/luns", params, &resp)
}
