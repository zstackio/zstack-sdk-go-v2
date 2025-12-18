// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorTrigger queries MonitorTrigger list
func (cli *ZSClient) QueryMonitorTrigger(params param.QueryParam) ([]view.MonitorTriggerInventoryView, error) {
	var resp []view.MonitorTriggerInventoryView
	return resp, cli.List("v1/monitoring/triggers", &params, &resp)
}
