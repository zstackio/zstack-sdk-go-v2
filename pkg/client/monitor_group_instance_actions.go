// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroupInstance 查询MonitorGroupInstance列表
func (cli *ZSClient) QueryMonitorGroupInstance(params param.QueryParam) ([]view.QueryMonitorGroupInstanceView, error) {
	var resp []view.QueryMonitorGroupInstanceView
	return resp, cli.List("v1/zwatch/monitorgroups/instances", &params, &resp)
}

