// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroupTemplateRef 查询MonitorGroupTemplateRef列表
func (cli *ZSClient) QueryMonitorGroupTemplateRef(params param.QueryParam) ([]view.QueryMonitorGroupTemplateRefView, error) {
	var resp []view.QueryMonitorGroupTemplateRefView
	return resp, cli.List("v1/zwatch/monitorgroups/monitortemplates/refs", &params, &resp)
}

