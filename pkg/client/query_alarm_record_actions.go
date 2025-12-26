// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAlarmRecord queries AlarmRecord list
func (cli *ZSClient) QueryAlarmRecord(params *param.QueryParam) ([]view.AlarmRecordsInventoryView, error) {
	var resp []view.AlarmRecordsInventoryView
	return resp, cli.List("v1/zwatch/alarm-records", params, &resp)
}
