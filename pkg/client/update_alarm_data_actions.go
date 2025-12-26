// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAlarmData updates AlarmData
func (cli *ZSClient) UpdateAlarmData(uuid string, params param.UpdateAlarmDataParam) (*view.UpdateAlarmDataEventView, error) {
	resp := view.UpdateAlarmDataEventView{}
	if err := cli.Put("v1/zwatch/alarm-histories/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
