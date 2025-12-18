// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeAlarmState 操作AlarmState
func (cli *ZSClient) ChangeAlarmState(uuid string, params param.ChangeAlarmStateParam) (*view.ChangeAlarmStateEventView, error) {
	resp := view.ChangeAlarmStateEventView{}
	if err := cli.Put("v1/zwatch/alarms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

