// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddLabelToAlarm 操作AddLabelToAlarm
func (cli *ZSClient) AddLabelToAlarm(params param.AddLabelToAlarmParam) (*view.AddLabelToAlarmEventView, error) {
	resp := view.AddLabelToAlarmEventView{}
	if err := cli.Post("v1/zwatch/alarms/{alarmUuid}/labels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

