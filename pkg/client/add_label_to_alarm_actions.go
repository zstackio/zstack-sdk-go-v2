// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddLabelToAlarm adds LabelToAlarm
func (cli *ZSClient) AddLabelToAlarm(params param.AddLabelToAlarmParam) (*view.AddLabelToAlarmEventView, error) {
	resp := view.AddLabelToAlarmEventView{}
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/labels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
