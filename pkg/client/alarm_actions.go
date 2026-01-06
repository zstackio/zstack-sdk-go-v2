// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlarm updates Alarm
func (cli *ZSClient) UpdateAlarm(uuid string, params param.UpdateAlarmParam) (*view.AlarmInventoryView, error) {
	var resp view.UpdateAlarmEventView
	if err := cli.Put("v1/zwatch/alarms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAlarm deletes Alarm
func (cli *ZSClient) DeleteAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms/{uuid}", uuid, string(deleteMode))
}
// QueryAlarm queries Alarm list
func (cli *ZSClient) QueryAlarm(params *param.QueryParam) ([]view.AlarmInventoryView, error) {
	var resp []view.AlarmInventoryView
	return resp, cli.List("v1/zwatch/alarms", params, &resp)
}
// CreateAlarm creates Alarm
func (cli *ZSClient) CreateAlarm(params param.CreateAlarmParam) (*view.AlarmInventoryView, error) {
	var resp view.CreateAlarmEventView
	if err := cli.Post("v1/zwatch/alarms", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
