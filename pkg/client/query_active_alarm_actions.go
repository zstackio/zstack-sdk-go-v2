// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryActiveAlarm queries ActiveAlarm list
func (cli *ZSClient) QueryActiveAlarm(params param.QueryParam) ([]view.ActiveAlarmInventoryView, error) {
	var resp []view.ActiveAlarmInventoryView
	return resp, cli.List("v1/zwatch/activealarms/alarms", &params, &resp)
}
