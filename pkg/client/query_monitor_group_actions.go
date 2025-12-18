// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroup queries MonitorGroup list
func (cli *ZSClient) QueryMonitorGroup(params param.QueryParam) ([]view.MonitorGroupInventoryView, error) {
	var resp []view.MonitorGroupInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups", &params, &resp)
}
