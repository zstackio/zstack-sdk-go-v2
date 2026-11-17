// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAlarm updates Alarm
func (cli *ZSClient) UpdateAlarm(uuid string, params param.UpdateAlarmParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/alarms", uuid, "", map[string]interface{}{
		"updateAlarm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAlarm deletes Alarm
func (cli *ZSClient) DeleteAlarm(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/alarms", uuid, string(deleteMode))
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

// PageAlarm Pagination
func (cli *ZSClient) PageAlarm(params *param.QueryParam) ([]view.AlarmInventoryView, int, error) {
	var alarms []view.AlarmInventoryView
	total, err := cli.Page("v1/zwatch/alarms", params, &alarms)
	return alarms, total, err
}
// CreateAlarm creates Alarm
func (cli *ZSClient) CreateAlarm(params param.CreateAlarmParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.Post("v1/zwatch/alarms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAlarmData updates AlarmData
func (cli *ZSClient) UpdateAlarmData(params param.UpdateAlarmDataParam) (*view.UpdateAlarmDataEventView, error) {
	resp := view.UpdateAlarmDataEventView{}
	if err := cli.PutWithRespKey("v1/zwatch/alarm-histories/actions", "", "", map[string]interface{}{
		"updateAlarmData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveActionFromAlarm removes ActionFromAlarm
func (cli *ZSClient) RemoveActionFromAlarm(alarmUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zwatch/alarms", alarmUuid, fmt.Sprintf("actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeAlarmState changes AlarmState
func (cli *ZSClient) ChangeAlarmState(uuid string, params param.ChangeAlarmStateParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/alarms", uuid, "", map[string]interface{}{
		"changeAlarmState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZWatchAlertHistogram gets ZWatchAlertHistogram by uuid
func (cli *ZSClient) GetZWatchAlertHistogram() (*view.GetZWatchAlertHistogramView, error) {
	var resp view.GetZWatchAlertHistogramView
	if err := cli.GetWithRespKey("v1/zwatch/alert-histories/histogram", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAlarmRecord queries AlarmRecord list
func (cli *ZSClient) QueryAlarmRecord(params *param.QueryParam) ([]view.AlarmRecordsInventoryView, error) {
	var resp []view.AlarmRecordsInventoryView
	return resp, cli.List("v1/zwatch/alarm-records", params, &resp)
}

func (cli *ZSClient) GetAlarmRecord(uuid string) (*view.AlarmRecordsInventoryView, error) {
	var resp view.AlarmRecordsInventoryView
	if err := cli.Get("v1/zwatch/alarm-records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAlarmRecord Pagination
func (cli *ZSClient) PageAlarmRecord(params *param.QueryParam) ([]view.AlarmRecordsInventoryView, int, error) {
	var alarmRecords []view.AlarmRecordsInventoryView
	total, err := cli.Page("v1/zwatch/alarm-records", params, &alarmRecords)
	return alarmRecords, total, err
}

// GetAlarmData gets AlarmData by uuid
func (cli *ZSClient) GetAlarmData() (*view.GetAlarmDataView, error) {
	var resp view.GetAlarmDataView
	if err := cli.GetWithRespKey("v1/zwatch/alarm-histories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddActionToAlarm adds ActionToAlarm
func (cli *ZSClient) AddActionToAlarm(alarmUuid string, params param.AddActionToAlarmParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zwatch/alarms/%s/actions", alarmUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
