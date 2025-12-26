// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMonitorGroupInstance queries MonitorGroupInstance list
func (cli *ZSClient) QueryMonitorGroupInstance(params *param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, error) {
	var resp []view.MonitorGroupInstanceInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/instances", params, &resp)
}
