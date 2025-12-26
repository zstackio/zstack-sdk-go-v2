// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddActionToAlarm adds ActionToAlarm
func (cli *ZSClient) AddActionToAlarm(params param.AddActionToAlarmParam) (*view.AddActionToAlarmEventView, error) {
	resp := view.AddActionToAlarmEventView{}
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
