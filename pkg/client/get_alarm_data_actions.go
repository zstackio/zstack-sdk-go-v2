// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAlarmData gets AlarmData by uuid
func (cli *ZSClient) GetAlarmData(uuid string) (*view.GetAlarmDataView, error) {
	var resp view.GetAlarmDataView
	if err := cli.Get("v1/zwatch/alarm-histories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
