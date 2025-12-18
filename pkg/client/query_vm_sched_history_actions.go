// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmSchedHistory queries VmSchedHistory list
func (cli *ZSClient) QueryVmSchedHistory(params param.QueryParam) ([]view.VmSchedHistoryInventoryView, error) {
	var resp []view.VmSchedHistoryInventoryView
	return resp, cli.List("v1/vm/sched-history", &params, &resp)
}
