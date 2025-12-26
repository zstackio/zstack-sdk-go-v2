// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetActiveAlarmStatus gets ActiveAlarmStatus by uuid
func (cli *ZSClient) GetActiveAlarmStatus(uuid string) (*view.GetActiveAlarmStatusView, error) {
	var resp view.GetActiveAlarmStatusView
	if err := cli.Get("v1/zwatch/activealarms/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
