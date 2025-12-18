// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryActiveAlarm 查询ActiveAlarm列表
func (cli *ZSClient) QueryActiveAlarm(params param.QueryParam) ([]view.QueryActiveAlarmView, error) {
	var resp []view.QueryActiveAlarmView
	return resp, cli.List("v1/zwatch/activealarms/alarms", &params, &resp)
}

