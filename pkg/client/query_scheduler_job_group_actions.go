// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySchedulerJobGroup queries SchedulerJobGroup list
func (cli *ZSClient) QuerySchedulerJobGroup(params *param.QueryParam) ([]view.SchedulerJobGroupInventoryView, error) {
	var resp []view.SchedulerJobGroupInventoryView
	return resp, cli.List("v1/scheduler/jobgroups", params, &resp)
}
