// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmSchedHistory queries VmSchedHistory list
func (cli *ZSClient) QueryVmSchedHistory(params *param.QueryParam) ([]view.VmSchedHistoryInventoryView, error) {
	var resp []view.VmSchedHistoryInventoryView
	return resp, cli.List("v1/vm/sched-history", params, &resp)
}
