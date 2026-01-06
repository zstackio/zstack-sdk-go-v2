// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupAlarm queries MonitorGroupAlarm list
func (cli *ZSClient) QueryMonitorGroupAlarm(params *param.QueryParam) ([]view.MonitorGroupAlarmInventoryView, error) {
	var resp []view.MonitorGroupAlarmInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/alarms", params, &resp)
}
