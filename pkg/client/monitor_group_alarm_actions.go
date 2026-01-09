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

func (cli *ZSClient) GetMonitorGroupAlarm(uuid string) (*view.MonitorGroupAlarmInventoryView, error) {
	var resp view.MonitorGroupAlarmInventoryView
	if err := cli.Get("v1/zwatch/monitorgroups/alarms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
