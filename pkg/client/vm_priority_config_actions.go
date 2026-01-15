// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmPriorityConfig queries VmPriorityConfig list
func (cli *ZSClient) QueryVmPriorityConfig(params *param.QueryParam) ([]view.VmPriorityConfigInventoryView, error) {
	var resp []view.VmPriorityConfigInventoryView
	return resp, cli.List("v1/vm-priority-config", params, &resp)
}

// PageVmPriorityConfig Pagination
func (cli *ZSClient) PageVmPriorityConfig(params *param.QueryParam) ([]view.VmPriorityConfigInventoryView, int, error) {
	var vmPriorityConfigs []view.VmPriorityConfigInventoryView
	total, err := cli.Page("v1/vm-priority-config", params, &vmPriorityConfigs)
	return vmPriorityConfigs, total, err
}
