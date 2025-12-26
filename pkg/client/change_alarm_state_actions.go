// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAlarmState changes AlarmState
func (cli *ZSClient) ChangeAlarmState(uuid string, params param.ChangeAlarmStateParam) (*view.ChangeAlarmStateEventView, error) {
	resp := view.ChangeAlarmStateEventView{}
	if err := cli.Put("v1/zwatch/alarms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
