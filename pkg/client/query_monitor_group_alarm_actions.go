// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroupAlarm queries MonitorGroupAlarm list
func (cli *ZSClient) QueryMonitorGroupAlarm(params param.QueryParam) ([]view.MonitorGroupAlarmInventoryView, error) {
	var resp []view.MonitorGroupAlarmInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/alarms", &params, &resp)
}
