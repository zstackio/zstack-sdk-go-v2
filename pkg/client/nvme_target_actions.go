// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNvmeTarget queries NvmeTarget list
func (cli *ZSClient) QueryNvmeTarget(params *param.QueryParam) ([]view.NvmeTargetInventoryView, error) {
	var resp []view.NvmeTargetInventoryView
	return resp, cli.List("v1/storage-devices/nvme/controllers", params, &resp)
}
// RefreshNvmeTarget operates on NvmeTarget
func (cli *ZSClient) RefreshNvmeTarget(params param.RefreshNvmeTargetParam) (*view.NvmeTargetInventoryView, error) {
	resp := view.NvmeTargetInventoryView{}
	if err := cli.Post("v1/storage-devices/nvme/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
