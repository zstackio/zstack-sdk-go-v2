// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddActionToAlarm 操作AddActionToAlarm
func (cli *ZSClient) AddActionToAlarm(params param.AddActionToAlarmParam) (*view.AddActionToAlarmEventView, error) {
	resp := view.AddActionToAlarmEventView{}
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

