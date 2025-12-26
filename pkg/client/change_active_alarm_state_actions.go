// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeActiveAlarmState changes ActiveAlarmState
func (cli *ZSClient) ChangeActiveAlarmState(uuid string, params param.ChangeActiveAlarmStateParam) (*view.ChangeActiveAlarmStateEventView, error) {
	resp := view.ChangeActiveAlarmStateEventView{}
	if err := cli.Put("v1/zwatch/activealarms/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
