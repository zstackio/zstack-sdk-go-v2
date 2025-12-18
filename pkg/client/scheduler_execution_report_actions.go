// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSchedulerExecutionReport 获取SchedulerExecutionReport详情
func (cli *ZSClient) GetSchedulerExecutionReport(uuid string) (*view.GetSchedulerExecutionReportView, error) {
	var resp view.GetSchedulerExecutionReportView
	if err := cli.Get("v1/scheduler/report", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

