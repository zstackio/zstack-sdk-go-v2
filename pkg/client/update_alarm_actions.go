// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAlarm updates Alarm
func (cli *ZSClient) UpdateAlarm(uuid string, params param.UpdateAlarmParam) (*view.UpdateAlarmEventView, error) {
	resp := view.UpdateAlarmEventView{}
	if err := cli.Put("v1/zwatch/alarms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
