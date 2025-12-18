// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAlarmData 更新AlarmData
func (cli *ZSClient) UpdateAlarmData(uuid string, params param.UpdateAlarmDataParam) (*view.UpdateAlarmDataEventView, error) {
	resp := view.UpdateAlarmDataEventView{}
	if err := cli.Put("v1/zwatch/alarm-histories/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

