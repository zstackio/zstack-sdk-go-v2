// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AckAlarmData operates on AckAlarmData
func (cli *ZSClient) AckAlarmData(params param.AckAlarmDataParam) (*view.AckAlertDataEventView, error) {
	resp := view.AckAlertDataEventView{}
	if err := cli.Post("v1/zwatch/alarm-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
