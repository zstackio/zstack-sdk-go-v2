// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySchedulerJobHistory queries SchedulerJobHistory list
func (cli *ZSClient) QuerySchedulerJobHistory(params *param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, error) {
	var resp []view.SchedulerJobHistoryInventoryView
	return resp, cli.List("v1/scheduler/job/history", params, &resp)
}
