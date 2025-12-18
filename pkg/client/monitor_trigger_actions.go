// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorTrigger 查询MonitorTrigger列表
func (cli *ZSClient) QueryMonitorTrigger(params param.QueryParam) ([]view.QueryMonitorTriggerView, error) {
	var resp []view.QueryMonitorTriggerView
	return resp, cli.List("v1/monitoring/triggers", &params, &resp)
}

