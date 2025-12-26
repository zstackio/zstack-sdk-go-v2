// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMonitorTemplate queries MonitorTemplate list
func (cli *ZSClient) QueryMonitorTemplate(params *param.QueryParam) ([]view.MonitorTemplateInventoryView, error) {
	var resp []view.MonitorTemplateInventoryView
	return resp, cli.List("v1/zwatch/monitortemplates", params, &resp)
}
