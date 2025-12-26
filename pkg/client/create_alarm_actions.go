// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAlarm creates Alarm
func (cli *ZSClient) CreateAlarm(params param.CreateAlarmParam) (*view.CreateAlarmEventView, error) {
	resp := view.CreateAlarmEventView{}
	if err := cli.Post("v1/zwatch/alarms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
