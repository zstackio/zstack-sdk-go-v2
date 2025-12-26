// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryMonitorTrigger queries MonitorTrigger list
func (cli *ZSClient) QueryMonitorTrigger(params *param.QueryParam) ([]view.MonitorTriggerInventoryView, error) {
	var resp []view.MonitorTriggerInventoryView
	return resp, cli.List("v1/monitoring/triggers", params, &resp)
}
