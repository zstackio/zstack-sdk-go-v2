// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAlarmLabel 更新AlarmLabel
func (cli *ZSClient) UpdateAlarmLabel(uuid string, params param.UpdateAlarmLabelParam) (*view.UpdateAlarmLabelEventView, error) {
	resp := view.UpdateAlarmLabelEventView{}
	if err := cli.Put("v1/zwatch/alarms/labels/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

