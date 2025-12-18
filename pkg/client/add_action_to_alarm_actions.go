// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddActionToAlarm adds ActionToAlarm
func (cli *ZSClient) AddActionToAlarm(params param.AddActionToAlarmParam) (*view.AddActionToAlarmEventView, error) {
	resp := view.AddActionToAlarmEventView{}
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
