// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmSchedHistory queries VmSchedHistory list
func (cli *ZSClient) QueryVmSchedHistory(params *param.QueryParam) ([]view.VmSchedHistoryInventoryView, error) {
	var resp []view.VmSchedHistoryInventoryView
	return resp, cli.List("v1/vm/sched-history", params, &resp)
}

func (cli *ZSClient) GetVmSchedHistory(uuid string) (*view.VmSchedHistoryInventoryView, error) {
	var resp view.VmSchedHistoryInventoryView
	if err := cli.Get("v1/vm/sched-history", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmSchedHistory Pagination
func (cli *ZSClient) PageVmSchedHistory(params *param.QueryParam) ([]view.VmSchedHistoryInventoryView, int, error) {
	var vmSchedHistories []view.VmSchedHistoryInventoryView
	total, err := cli.Page("v1/vm/sched-history", params, &vmSchedHistories)
	return vmSchedHistories, total, err
}
