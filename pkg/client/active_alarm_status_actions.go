// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetActiveAlarmStatus 获取ActiveAlarmStatus详情
func (cli *ZSClient) GetActiveAlarmStatus(uuid string) (*view.GetActiveAlarmStatusView, error) {
	var resp view.GetActiveAlarmStatusView
	if err := cli.Get("v1/zwatch/activealarms/status", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

