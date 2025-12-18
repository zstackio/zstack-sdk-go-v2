// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroupAlarm 查询MonitorGroupAlarm列表
func (cli *ZSClient) QueryMonitorGroupAlarm(params param.QueryParam) ([]view.QueryMonitorGroupAlarmView, error) {
	var resp []view.QueryMonitorGroupAlarmView
	return resp, cli.List("v1/zwatch/monitorgroups/alarms", &params, &resp)
}

