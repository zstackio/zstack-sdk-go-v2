// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAlarmRecord 查询AlarmRecord列表
func (cli *ZSClient) QueryAlarmRecord(params param.QueryParam) ([]view.QueryAlarmRecordView, error) {
	var resp []view.QueryAlarmRecordView
	return resp, cli.List("v1/zwatch/alarm-records", &params, &resp)
}

