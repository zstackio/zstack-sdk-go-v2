// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryActiveAlarm queries ActiveAlarm list
func (cli *ZSClient) QueryActiveAlarm(params *param.QueryParam) ([]view.ActiveAlarmInventoryView, error) {
	var resp []view.ActiveAlarmInventoryView
	return resp, cli.List("v1/zwatch/activealarms/alarms", params, &resp)
}
