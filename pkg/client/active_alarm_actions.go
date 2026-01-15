// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryActiveAlarm queries ActiveAlarm list
func (cli *ZSClient) QueryActiveAlarm(params *param.QueryParam) ([]view.ActiveAlarmInventoryView, error) {
	var resp []view.ActiveAlarmInventoryView
	return resp, cli.List("v1/zwatch/activealarms/alarms", params, &resp)
}

func (cli *ZSClient) GetActiveAlarm(uuid string) (*view.ActiveAlarmInventoryView, error) {
	var resp view.ActiveAlarmInventoryView
	if err := cli.Get("v1/zwatch/activealarms/alarms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageActiveAlarm Pagination
func (cli *ZSClient) PageActiveAlarm(params *param.QueryParam) ([]view.ActiveAlarmInventoryView, int, error) {
	var activeAlarms []view.ActiveAlarmInventoryView
	total, err := cli.Page("v1/zwatch/activealarms/alarms", params, &activeAlarms)
	return activeAlarms, total, err
}
