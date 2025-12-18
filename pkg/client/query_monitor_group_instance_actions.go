// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMonitorGroupInstance queries MonitorGroupInstance list
func (cli *ZSClient) QueryMonitorGroupInstance(params param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, error) {
	var resp []view.MonitorGroupInstanceInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/instances", &params, &resp)
}
