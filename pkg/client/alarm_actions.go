// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlarm updates Alarm
func (cli *ZSClient) UpdateAlarm(uuid string, params param.UpdateAlarmParam) (*view.AlarmInventoryView, error) {
	var resp view.UpdateAlarmEventView
	err := cli.PutWithSpec("v1/zwatch/alarms", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAlarm deletes Alarm
func (cli *ZSClient) DeleteAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zwatch/alarms", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryAlarm queries Alarm list
func (cli *ZSClient) QueryAlarm(params *param.QueryParam) ([]view.AlarmInventoryView, error) {
	var resp []view.AlarmInventoryView
	return resp, cli.List("v1/zwatch/alarms", params, &resp)
}

func (cli *ZSClient) GetAlarm(uuid string) (*view.AlarmInventoryView, error) {
	var resp view.AlarmInventoryView
	if err := cli.Get("v1/zwatch/alarms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAlarm creates Alarm
func (cli *ZSClient) CreateAlarm(params param.CreateAlarmParam) (*view.AlarmInventoryView, error) {
	var resp view.CreateAlarmEventView
	if err := cli.Post("v1/zwatch/alarms", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
