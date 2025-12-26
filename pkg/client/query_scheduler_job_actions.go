// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySchedulerJob queries SchedulerJob list
func (cli *ZSClient) QuerySchedulerJob(params *param.QueryParam) ([]view.SchedulerJobInventoryView, error) {
	var resp []view.SchedulerJobInventoryView
	return resp, cli.List("v1/scheduler/jobs", params, &resp)
}
