// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupAlarm queries MonitorGroupAlarm list
func (cli *ZSClient) QueryMonitorGroupAlarm(params *param.QueryParam) ([]view.MonitorGroupAlarmInventoryView, error) {
	var resp []view.MonitorGroupAlarmInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/alarms", params, &resp)
}

// PageMonitorGroupAlarm Pagination
func (cli *ZSClient) PageMonitorGroupAlarm(params *param.QueryParam) ([]view.MonitorGroupAlarmInventoryView, int, error) {
	var monitorGroupAlarms []view.MonitorGroupAlarmInventoryView
	total, err := cli.Page("v1/zwatch/monitorgroups/alarms", params, &monitorGroupAlarms)
	return monitorGroupAlarms, total, err
}
