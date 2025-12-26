// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMonitorGroupTemplateRef queries MonitorGroupTemplateRef list
func (cli *ZSClient) QueryMonitorGroupTemplateRef(params *param.QueryParam) ([]view.MonitorGroupTemplateRefInventoryView, error) {
	var resp []view.MonitorGroupTemplateRefInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/monitortemplates/refs", params, &resp)
}
