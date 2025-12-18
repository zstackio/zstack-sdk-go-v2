// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySchedulerJobHistory queries SchedulerJobHistory list
func (cli *ZSClient) QuerySchedulerJobHistory(params param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, error) {
	var resp []view.SchedulerJobHistoryInventoryView
	return resp, cli.List("v1/scheduler/job/history", &params, &resp)
}
