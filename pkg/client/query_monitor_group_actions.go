// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMonitorGroup queries MonitorGroup list
func (cli *ZSClient) QueryMonitorGroup(params *param.QueryParam) ([]view.MonitorGroupInventoryView, error) {
	var resp []view.MonitorGroupInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups", params, &resp)
}
