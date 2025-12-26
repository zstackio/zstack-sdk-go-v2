// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAlarm queries Alarm list
func (cli *ZSClient) QueryAlarm(params *param.QueryParam) ([]view.AlarmInventoryView, error) {
	var resp []view.AlarmInventoryView
	return resp, cli.List("v1/zwatch/alarms", params, &resp)
}
