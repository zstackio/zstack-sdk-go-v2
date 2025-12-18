// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroup 查询MonitorGroup列表
func (cli *ZSClient) QueryMonitorGroup(params param.QueryParam) ([]view.QueryMonitorGroupView, error) {
	var resp []view.QueryMonitorGroupView
	return resp, cli.List("v1/zwatch/monitorgroups", &params, &resp)
}

