// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySchedulerJobHistory 查询SchedulerJobHistory列表
func (cli *ZSClient) QuerySchedulerJobHistory(params param.QueryParam) ([]view.QuerySchedulerJobHistoryView, error) {
	var resp []view.QuerySchedulerJobHistoryView
	return resp, cli.List("v1/scheduler/job/history", &params, &resp)
}

