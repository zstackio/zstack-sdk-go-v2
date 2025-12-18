// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAlarm 更新Alarm
func (cli *ZSClient) UpdateAlarm(uuid string, params param.UpdateAlarmParam) (*view.UpdateAlarmEventView, error) {
	resp := view.UpdateAlarmEventView{}
	if err := cli.Put("v1/zwatch/alarms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

