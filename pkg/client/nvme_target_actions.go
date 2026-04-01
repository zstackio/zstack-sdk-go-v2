// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNvmeTarget queries NvmeTarget list
func (cli *ZSClient) QueryNvmeTarget(params *param.QueryParam) ([]view.NvmeTargetInventoryView, error) {
	var resp []view.NvmeTargetInventoryView
	return resp, cli.List("v1/storage-devices/nvme/controllers", params, &resp)
}

func (cli *ZSClient) GetNvmeTarget(uuid string) (*view.NvmeTargetInventoryView, error) {
	var resp view.NvmeTargetInventoryView
	if err := cli.Get("v1/storage-devices/nvme/controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNvmeTarget Pagination
func (cli *ZSClient) PageNvmeTarget(params *param.QueryParam) ([]view.NvmeTargetInventoryView, int, error) {
	var nvmeTargets []view.NvmeTargetInventoryView
	total, err := cli.Page("v1/storage-devices/nvme/controllers", params, &nvmeTargets)
	return nvmeTargets, total, err
}
// RefreshNvmeTarget operates on NvmeTarget
func (cli *ZSClient) RefreshNvmeTarget(params param.RefreshNvmeTargetParam) (*view.NvmeTargetInventoryView, error) {
	resp := view.NvmeTargetInventoryView{}
	if err := cli.Post("v1/storage-devices/nvme/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
