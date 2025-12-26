// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMonitorTriggerAction queries MonitorTriggerAction list
func (cli *ZSClient) QueryMonitorTriggerAction(params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, error) {
	var resp []view.MonitorTriggerActionInventoryView
	return resp, cli.List("v1/monitoring/trigger-actions", params, &resp)
}
